package table

import (
	"github.com/shopspring/decimal"
)

type CoinConfig struct {
	Id           int64
	CoinName     string
	Hash         string
	Height       int64
	BlockTime    string
	Ctime        string
	Mtime        string
}

type ApiAccount struct {
	Id             int64
	name           string
	role           string
	token          string
	Ctime          string
	Mtime          string
}
