package db

import (
	"database/sql"
	"fmt"
	"sync_blockchain/db/table"
)

type ApiAccountModel struct{
	model
}

func NewApiConfigModel(stmt *sql.Tx) *ApiConfigModel{
	m := &ApiConfigModel{
		model{
			//TODO::顺序必须与 rows.Scan 和 this.Insert 的参数一致
			fields:"`id`,`name`,`role`,`token`, `ctime`,`mtime`",
			table:"`api_config`",
			scan: func(rows *sql.Rows) (interface{}, error) {
				var (
					list []*table.ApiAccount
					//row table.Addresses
					scanErr   error
				)

				for rows.Next() {
					var row table.ApiAccount
					scanErr = rows.Scan(&row.Id, &row.Name, &row.Role, &row.Token, &row.Ctime, &row.Mtime)
					if scanErr != nil {
						return nil, scanErr
					}
					list = append(list, &row)
				}
				return list, nil
			},
		},
	}
	if stmt == nil{
		m.stmt = db
	}else{
		m.stmt = stmt
	}
	return m
}

func (this *ApiAccountModel) Insert(addr *table.ApiAccount) int64{
	return this.model.Insert(nil, addr.Name, addr.Role, addr.Token, addr.Ctime, addr.Mtime)
}

func (this *ApiAccountModel)Select(where string, order string, limit string)(bool, []*table.ApiAccount){
	ok, addrs := this.model.Select(where, order, limit)
	if addrs == nil{
		return ok, nil
	}
	return ok, addrs.([]*table.ApiAccount)
}

func (this *ApiAccountModel) FindByName(name string) (bool, *table.ApiAccount){
	ok, row := this.Select(" `name`='" + name + "' ", "", "1")
	if row == nil{
		return ok, nil
	}
	return ok, row[0]
}

func (this *ApiAccountModel) FindByToken(token string) (bool, *table.ApiAccount){
	ok, row := this.Select(" `token`='" + token + "' ", "", "1")
	if row == nil{
		return ok, nil
	}
	return ok, row[0]
}


func (this *ApiAccountModel) SelectAll()(bool, []*table.ApiAccount){
	return this.Select("", "", "")
}


func (this *ApiAccountModel) FindById(id int64) (bool, *table.ApiAccount){
	ok, api_configs := this.Select(fmt.Sprintf("`id`='%d'", id), "", "1")
	if api_configs == nil{
		return ok, nil
	}
	return ok, api_configs[0]
}

func (this *ApiAccountModel) Exists(name string) (bool, bool){
	ok, count := this.Count(" `name`='" + name + "' ")
	return ok, count > 0
}
