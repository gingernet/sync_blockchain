package main

import (
	"fmt"
	"strconv"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

var (
	client *rpc.Client
)

func init() {
	client, _ = rpc.Dial("139.180.209.214:7654")
	if client == nil {
		fmt.Println("rpc.Dial err")
		return
	}
}

//获取账户列表
func getAccounts(client *rpc.Client) (accounts []string, err error) {
	err = client.Call(&accounts, "eth_accounts")
	if err == nil {
		return accounts, nil
	} else {
		return nil, errors.New("账户列表获取错误")
	}

}

//获取挖矿账户
func getCoinbase(client *rpc.Client) (coinbase string, err error) {
	err = client.Call(&coinbase, "eth_coinbase")
	if err == nil {
		return coinbase, nil
	} else {
		return "", errors.New("挖矿账户获取错误")
	}
}

//获取余额
func getBalance(client *rpc.Client, account string) (Balance int64, err error) {

	var balance string
	err = client.Call(&balance, "eth_getBalance", account, "latest")
	if err != nil {
		return -1, err
	}
	Balance, _ = strconv.ParseInt(balance, 0, 64)
	return Balance, nil

}
func creatNewAccount(client *rpc.Client, password string) (newAccount string, err error) {
	err = client.Call(&newAccount, "personal_newAccount", password)
	if err != nil {
		return "", err
	}
	return newAccount, nil

}
func main() {

	//创建新账户
	var password string = "123456"
	newAccount, err := creatNewAccount(client, password)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("新账户为：", newAccount)

	//获取账户列表
	accounts, err := getAccounts(client)
	if err != nil {
		fmt.Println("err=", err)
	}
	for i, v := range accounts {
		balance, err := getBalance(client, v)
		if err != nil {
			fmt.Println("err=", err)
		} else {
			fmt.Printf("账户%d的账号为：%s，余额为：%d\n", i, v, balance)
		}

	}

	//获取挖矿账户
	coinbase, err := getCoinbase(client)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("挖矿账户为：", coinbase)

	//延迟关闭
	defer client.Close()

}
