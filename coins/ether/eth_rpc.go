package ether

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
func GetAccounts(client *rpc.Client) (accounts []string, err error) {
	err = client.Call(&accounts, "eth_accounts")
	if err == nil {
		return accounts, nil
	} else {
		return nil, errors.New("账户列表获取错误")
	}

}

//获取挖矿账户
func GetCoinbase(client *rpc.Client) (coinbase string, err error) {
	err = client.Call(&coinbase, "eth_coinbase")
	if err == nil {
		return coinbase, nil
	} else {
		return "", errors.New("挖矿账户获取错误")
	}
}

//获取余额
func GetBalance(client *rpc.Client, account string) (Balance int64, err error) {

	var balance string
	err = client.Call(&balance, "eth_getBalance", account, "latest")
	if err != nil {
		return -1, err
	}
	Balance, _ = strconv.ParseInt(balance, 0, 64)
	return Balance, nil
}


func CreatNewAccount(client *rpc.Client, password string) (newAccount string, err error) {
	err = client.Call(&newAccount, "personal_newAccount", password)
	if err != nil {
		return "", err
	}
	return newAccount, nil
}

