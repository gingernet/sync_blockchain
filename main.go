package main

import (
	"fmt"
	"os"
	"sync_blockchain/config"
	"sync_blockchain/db"
)


func initConfig()  {
	cfg = config.New(cfgFile, "./resource/app.yaml")
	if cfg == nil {
		fmt.Errorf("Initialization configuration file exception, cfg = %+v", cfg)
		os.Exit(-1)
	}

	// DB  init
	//db.Init(cfg.Dsn, cfg.MaxConn, cfg.MaxIdle)
	//cfg.LoadCoinDataConf(auth)

	// 打印配置数据
	// fmt.Println("%+v", cfg)
}

func main()  {
	initConfig()
}
