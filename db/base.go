package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB  //连接池

//事务处理 返回error触发回滚 返回nil触发commit
type TransFn func(stmt *sql.Tx) error

//query rows scan
type rowsScan func(rows *sql.Rows)(interface{}, error)


type statement interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

/**
 * 初始化配置
 * @dsn mysql conn string
 * @maxconn db pool max conn
 * @maxidle db pool max idle
 */
func Init(dsn string, maxconn int, maxidle int) {
	var err error
	db, _ = sql.Open("mysql", dsn)
	db.SetMaxOpenConns(maxconn)
	db.SetMaxIdleConns(maxidle)
	//db.SetConnMaxLifetime(300)
	//test connection
	err = db.Ping()
	if err != nil{
		panic(err.Error())
	}
}

/***
 * 事务封装
 * @fn func 业务逻辑处理函数 返回error触发回滚 返回nil触发提交
 */
func WithTransaction(fn TransFn) (err error) {
	stmt, err := db.Begin()
	if err != nil {
		return
	}

	fmt.Println("transaction begin  ")
	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			stmt.Rollback()
			fmt.Errorf("transaction rollback with painc  ")
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			stmt.Rollback()
			fmt.Errorf("transaction rollback with err  ")
		} else {
			// all good, commit
			err = stmt.Commit()
			fmt.Println("transaction commit  ")
		}
		fmt.Println("transaction end \r\n")
	}()

	err = fn(stmt)
	return err
}

func findColumn(stmt statement, sqlStr string, args ...interface{})(ret string, err error){
	find(stmt, sqlStr, func(rows *sql.Rows) (interface{}, error) {
		if rows.Next() {
			err = rows.Scan(&ret)
		}
		return nil, nil
	}, args...)

	if err  == sql.ErrNoRows {
		err = nil
	}
	return
}

func count(stmt statement, sqlStr string, args ...interface{})(int64, error){
	res, err := find(stmt, sqlStr, func(rows *sql.Rows) (interface{}, error) {
		var (
			count int64
			scanErr   error
		)

		if rows.Next() {
			scanErr = rows.Scan(&count)
			if scanErr == nil || scanErr == sql.ErrNoRows {
				return count, nil
			}
			return 0, scanErr
		}
		return 0, nil
	}, args...)

	if err != nil {
		return 0, err
	}
	return res.(int64), nil
}

func find(stmt statement, sqlStr string, scan rowsScan, args ...interface{})(result interface{}, err error){
	var rows *sql.Rows
	rows, err = stmt.Query(sqlStr, args...)
	if err != nil{
		fmt.Println("sqlStr: %s, \nargs: %+v,  select err:%+v \n", sqlStr, args, err)
		return
	}
	defer rows.Close()
	result, err = scan(rows)
	return
}

func insert(stmt statement, sqlStr string, args ...interface{}) (int64, error) {
	res, err := stmt.Exec(sqlStr, args...)
	if err != nil{
		return 0, err
	}
	return res.LastInsertId()
}

func exec(stmt statement, sqlStr string, args ...interface{}) (int64, error) {
	res, err := stmt.Exec(sqlStr, args...)
	if err != nil{
		return 0, err
	}
	return res.RowsAffected()
}
