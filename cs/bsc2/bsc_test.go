package bsc2

import (
	"fmt"
	"github.com/362228416/go-eth/cs/eth2"
	"testing"
)

func TestGetBlockNumber(t *testing.T) {
	InitBscWeb3("dev")
	fmt.Println(eth2.EthGetBlockNumber())
	InitBscWeb3("test")
	fmt.Println(eth2.EthGetBlockNumber())
	InitBscWeb3("prod")
	fmt.Println(eth2.EthGetBlockNumber())
	InitBscWeb3("main")
	fmt.Println(eth2.EthGetBlockNumber())
}
