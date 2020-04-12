package main

import (
	"fmt"
	 "sync_blockchain/coins/ether"
)


func main()  {
	fmt.Println("start block chain sync")

	var password string = "123456"
	newAccount, err := ether.CreatNewAccount(client, password)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("新账户为：", newAccount)


	accounts, err := ether.GetAccounts(client)
	if err != nil {
		fmt.Println("err=", err)
	}
	for i, v := range accounts {
		balance, err := ether.GetBalance(client, v)
		if err != nil {
			fmt.Println("err=", err)
		} else {
			fmt.Printf("账户%d的账号为：%s，余额为：%d\n", i, v, balance)
		}

	}

	coinbase, err := ether.GetCoinbase(client)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("挖矿账户为：", coinbase)

	defer client.Close()
	fmt.Println("end block chain sync")
}
