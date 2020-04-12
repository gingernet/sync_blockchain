package db

import (
	"database/sql"
	"sync_blockchain/db/table"
)

type CoinConfigModel struct{
	model
}

func NewCoinConfigModelModel(stmt *sql.Tx) *CoinConfigModelModel{
	m := &CoinConfigModelModel{
		model{
			fields:"`id`,`coin_name`,`hash`,`height`, `block_time`, `ctime`,`mtime`",
			table:"`coin_config`",
			scan: func(rows *sql.Rows) (interface{}, error) {
				var (
					list []*table.CoinConfig
					scanErr   error
				)
				for rows.Next() {
					var row table.CoinConfig
					scanErr = rows.Scan(&row.Id, &row.CoinName, &row.Hash, &row.Height, &row.BlockTime, &row.Ctime, &row.Mtime)
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

func (this *CoinConfigModel) Insert(coin_config *table.CoinConfig) int64{
	return this.model.Insert(nil, coin_config.CoinName, coin_config.Hash, coin_config.Height, coin_config.BlockTime, coin_config.Ctime, coin_config.Mtime)
}

func (this *CoinConfigModel)Select(where string, order string, limit string)(bool, []*table.CoinConfig){
	ok, coin_configs := this.model.Select(where, order, limit)
	if coin_configs == nil{
		return ok, nil
	}
	return ok, coin_configs.([]*table.CoinConfig)
}

func (this *ApiAccountModel) FindByCoinName(coin_name string) (bool, *table.CoinConfig){
	ok, row := this.Select(" `coin_name`='" + coin_name + "' ", "", "1")
	if row == nil{
		return ok, nil
	}
	return ok, row[0]
}

func (this *CoinConfigModel) SelectAll()(bool, []*table.CoinConfig){
	return this.Select("", "", "")
}

func (this *CoinConfigModel) Exists(coin_name string) (bool, bool){
	ok, count := this.Count(" `coin_name`='" + coin_name + "' ")
	return ok, count > 0
}

