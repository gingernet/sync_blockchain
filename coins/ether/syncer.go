package ether

import (
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync_blockchain/config"
)

type BlockInfo struct {
	Number *big.Int
	Hash   string
}

type ChainSyncer struct {
	Conf *config.BaseConfig
	CurrentBlock *BlockInfo           		// 当前块信息
	LastBlockNumber *big.Int         		// 前一个块高的高度
	LastBlockTime *big.Int          		// 前一个块的时间
	RpcClients map[int]*ethclient.Client    // rpc map
}
