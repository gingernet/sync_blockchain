package db

import (
	"database/sql"
	"fmt"
	"sync_blockchain/db/table"
	"time"
	"strings"
)

type ApiAccountModel struct{
	model
}

func NewAddressModel(stmt *sql.Tx) *AddressesModel{
	m := &AddressesModel{
		model{
			//TODO::顺序必须与 rows.Scan 和 this.Insert 的参数一致
			fields:"`id`,`address`,`nonce`,`type`,`hash`,`ctime`,`mtime`",
			table:"`addresses`",
			scan: func(rows *sql.Rows) (interface{}, error) {
				var (
					list []*table.ApiAccount
					//row table.Addresses
					scanErr   error
				)

				for rows.Next() {
					var row table.Addresses
					scanErr = rows.Scan(&row.Id, &row.Address, &row.Nonce, &row.Type, &row.Hash, &row.Ctime, &row.Mtime)
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

func (this *AddressesModel) Insert(addr *table.Addresses) int64{
	return this.model.Insert(nil, addr.Address, addr.Nonce, addr.Type, addr.Hash, addr.Ctime, addr.Mtime)
}

func (this *AddressesModel)Select(where string, order string, limit string)(bool, []*table.Addresses){
	ok, addrs := this.model.Select(where, order, limit)
	if addrs == nil{
		return ok, nil
	}
	return ok, addrs.([]*table.Addresses)
}

func (this *AddressesModel) FindByAddr(addr string) (bool, *table.Addresses){
	ok, row := this.Select(" `address`='" + addr + "' ", "", "1")
	if row == nil{
		return ok, nil
	}
	return ok, row[0]
}

func (this *AddressesModel) FindByType(addrType int) (bool, *table.Addresses){
	ok, row := this.Select(fmt.Sprintf(" `type`='%d'", addrType), "", "1")
	if row == nil{
		return ok, nil
	}
	return ok, row[0]
}

func (this *AddressesModel) SelectByAddrs(addrs []string) (bool, []*table.Addresses){
	ok, rows := this.Select(" `address` in ('" +  strings.Join(addrs, "','") + "') ", "", "1")
	return ok, rows
}

func (this *AddressesModel) SelectAll()(bool, []*table.Addresses){
	return this.Select("", "", "")
}

func (this *AddressesModel) Page(pageIndex int, pageSize int)(bool, []*table.Addresses){
	offset := (pageIndex - 1) * pageSize
	return this.Select("", "", fmt.Sprintf("%d,%d", offset, pageSize))
}

func (this *AddressesModel) TotalCount()(bool, int64){
	return this.model.Count("")
}

func (this *AddressesModel) IncrNonce(id int64)int64{
	sqlStr := fmt.Sprintf("update %s set `nonce`=`nonce`+1, `mtime`=? where `id`='%d' limit 1", this.table, id)
	return this.model.Save(sqlStr, time.Now().Format("2006-01-02 15:04:05"))
}

func (this *AddressesModel) FindById(id int64) (bool, *table.Addresses){
	ok, addrs := this.Select(fmt.Sprintf("`id`='%d'", id), "", "1")
	if addrs == nil{
		return ok, nil
	}
	return ok, addrs[0]
}

func (this *AddressesModel) Exists(addr string) (bool, bool){
	ok, count := this.Count(" `address`='" + addr + "' ")
	return ok, count > 0
}

func (this *AddressesModel) FindByAddress(addr string) (bool, *table.Addresses){
	ok, addrs := this.Select(fmt.Sprintf("`address`='%s'", addr), "", "1")
	if addrs == nil{
		return ok, nil
	}
	return ok, addrs[0]
}

func (this *AddressesModel) IncrNonceByAddress(addr string)int64{
	sqlStr := fmt.Sprintf("update %s set `nonce`=`nonce`+1, `mtime`=? where `address`=? limit 1", this.table)
	return this.model.Save(sqlStr, time.Now().Format("2006-01-02 15:04:05"), addr)
}

func (this *AddressesModel) DecrNonceByAddress(addr string)int64{
	sqlStr := fmt.Sprintf("update %s set `nonce`=`nonce`-1, `mtime`=? where `address`=? limit 1", this.table)
	return this.model.Save(sqlStr, time.Now().Format("2006-01-02 15:04:05"), addr)
}
