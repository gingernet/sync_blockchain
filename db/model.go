package db

import (
	"fmt"
	"strings"
)

type model struct {
	fields string
	table string
	scan rowsScan
	stmt statement
}

func (this *model) Select(where string, order string, limit string) (bool, interface{}){
	sqlStr := "SELECT " + this.fields + " FROM  " + this.table
	if where != ""{
		sqlStr += " WHERE " + where
	}
	if order != ""{
		sqlStr += " ORDER BY  " + order
	}
	if limit != ""{
		sqlStr += " LIMIT " + limit
	}
	rows, err := find(this.stmt, sqlStr, this.scan)
	if err != nil{
		fmt.Errorf("sqlStr:%s select err:%+v \n", sqlStr, err)
		return false, nil
	}

	if rows == nil {
		fmt.Println("sqlStr:%s select err:%+v \n", sqlStr, " empty row ")
		return true, nil
	}
	return true, rows
}

func (this *model) Insert(args ...interface{}) int64{
	sqlStr := "INSERT INTO  " + this.table + " (" + this.fields + ")VALUES(" + strings.Repeat("?,", len(args) - 1) + "?)"
	id, err := insert(this.stmt, sqlStr, args...);
	if err != nil{
		fmt.Errorf("sqlStr: %s, \nargs: %+v,  insert err:%+v \n", sqlStr, args, err)
	}else{
		fmt.Errorf("sqlStr: %s, \nargs: %+v,  insert succ ID:%d \n", sqlStr, args, id)
	}
	return id
}

func (this *model) Save(sqlStr string, args ...interface{}) int64{
	rowEffect, err := exec(this.stmt, sqlStr, args...);
	if err != nil{
		fmt.Errorf.Errorf("sqlStr: %s, \nargs:%+v, update err:%+v \n", sqlStr, args, err)
	}else{
		fmt.Println("sqlStr: %s, \nargs:%+v, update succ row:%d \n", sqlStr, args, rowEffect)
	}
	return rowEffect
}

func (this *model) Delete(where string, limit string)int64{
	sqlStr := "delete from " + this.table
	if where != ""{
		sqlStr += " where " + where
	} else{
		panic("delete " + this.table + " without condition")
	}
	if limit != ""{
		sqlStr += " LIMIT " + limit
	}

	rowEffect, err := exec(this.stmt, sqlStr)
	if err != nil{
		fmt.Errorf("sqlStr: %s, del err:%+v \n", sqlStr, err)
	} else{
		fmt.Write().Debugf("sqlStr: %s, del succ row:%d \n", sqlStr, rowEffect)
	}
	return rowEffect
}

func (this *model) Count(where string) (bool, int64){
	sqlStr := "select count(*) from " + this.table
	if where != ""{
		sqlStr += " where " + where
	}
	count, err := count(this.stmt, sqlStr)
	if err != nil{
		fmt.Errorf("sqlStr: %s, count err:%+v \n", sqlStr, err)
		return false, 0
	} else{
		fmt.Println("error in func")
	}
	return true, count
}

func (this *model) Sum(field string, where string) (bool, string){
	sqlStr := "select sum("+field+") from " + this.table
	if where != ""{
		sqlStr += " where " + where
	}
	ret, err := findColumn(this.stmt, sqlStr)
	if err != nil{
		fmt.Errorf("sqlStr: %s, sum err:%+v \n", sqlStr, err)
		return false, ""
	} else{
		fmt.Println("sqlStr: %s, sum succ row:%s \n", sqlStr, ret)
	}
	return true, ret
}

