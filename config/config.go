package config

import (
	"fmt"
	"os"
	"math/big"
	"strings"
	"github.com/spf13/viper"
	// "github.com/shopspring/decimal"
)


// 配置文件字段申明
type BaseConfig struct {

	CoinType string          // 币种类型
	RpcUrl string           // 节点机链接
	RPCUrls []string        // ,分割 RpcUrl
	ChainID *big.Int        // 币种ID

	TokenABI   string           // ABI
	InitialHeight int64        	// 初始块高

	CoinsConf map[string] *CoinDataConf
	CoinsConfContractAddr map[string] *CoinDataConf

	// API
	ApiUrl string
	ApiId  string
	ApiSecret string

	// DB
	Dsn     string
	MaxConn int
	MaxIdle int

	// 钉钉
	DingDingBotURL   string
	DingToken 	string
	MsgType  	int64

	// ES
	EsUrl string
	ESId  string
	ESSecret string


	SafetyConfirmations int64       // 安全块高
}


// CoinData 配置字段申明
type CoinDataConf struct {
	Symbol string   			// 币种
	Coin_type int  				// 0：eth,1：ERC20，2：ERC223，3：ERC773
	Contract_addr string    	// 合约地址
	Precision int       		// 精度，正整数<=18
}


/**
 * 基础配置文件初始值
 *
 * @return *BaseConfig
 */
func DefaultBaseConfig () *BaseConfig {
	return &BaseConfig {
		CoinType : "eth, btc, omni-usdt",
		RpcUrl : "",
		ChainID : big.NewInt(0),
		TokenABI : "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"push\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
		InitialHeight : 0,
		SafetyConfirmations: 6,
	}
}


func (cfg *BaseConfig) LoadCoinDataConf(auth string)  {

}


func New (cfgFile string, strPath string) *BaseConfig {
	// 通过viper 类库, 解析 yaml 配置
	if cfgFile == "" {
		fmt.Println("yaml file is empty(配置文件不能为空)")
		os.Exit(-1)
	}

	viper.SetConfigFile(cfgFile)
	viper.SetConfigName("app")
	if strPath == "" {
		viper.AddConfigPath(".")
	} else {
		viper.AddConfigPath(strPath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can't read config:%+v, run with default config.")
		os.Exit(-1)
	}

	// 获取默认配置
	cfg := DefaultBaseConfig()

	// 节点机, node
	urls := strings.Split(getString("node.rpcUrl", cfg.RpcUrl), ",")
	cfg.RPCUrls = make([]string, len(urls))
	for i, v := range urls {
		cfg.RPCUrls[i] = strings.Replace(v, " ", "", -1)
	}

	// 签名机 Sign
	cfg.ChainID = big.NewInt(getInt64("sign.chainID", cfg.ChainID.Int64()))

	// Api
	cfg.ApiUrl = getString("api.url", cfg.ApiUrl)
	cfg.ApiId = getString("api.id", cfg.ApiId)
	cfg.ApiSecret = getString("api.secret", cfg.ApiSecret)

	// 币种类型
	cfg.CoinType = strings.ToLower(getString("wallet.coinType", cfg.CoinType))
	// 初始块高
	cfg.InitialHeight = getInt64("wallet.height", cfg.InitialHeight)

	// 安全块高
	cfg.SafetyConfirmations = getInt64("wallet.safetyConfirmations", cfg.SafetyConfirmations)

	// DB
	cfg.Dsn = getString("db.dsn", cfg.Dsn)
	cfg.MaxConn = getInt("db.maxConn", cfg.MaxConn)
	cfg.MaxIdle = getInt("db.maxIdle", cfg.MaxIdle)

	return cfg
}
