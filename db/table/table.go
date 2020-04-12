package table


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
	Name           string
	Role           string
	Token          string
	Ctime          string
	Mtime          string
}

