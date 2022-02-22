package bsc2

import (
	"github.com/362228416/go-eth/cs/eth2"
)

func InitBscWeb3(profile string) {
	var web3Url string
	var chainId int64 = 0x38
	if profile == "dev" || profile == "test" {
		//web3Url = "https://data-seed-prebsc-1-s1.binance.org:8545/"
		web3Url = "https://data-seed-prebsc-1-s2.binance.org:8545/"
		chainId = 0x61
	} else if profile == "prod" || profile == "main" {
		web3Url = "https://bsc-dataseed1.ninicoin.io"
	}
	eth2.InitWeb3(web3Url, chainId)
}
