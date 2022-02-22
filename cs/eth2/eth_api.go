package eth2

import "github.com/ethereum/go-ethereum/common/hexutil"

func EthGetBlockNumber() uint64 {
	var last string
	_rpc.Call(&last, "eth_blockNumber")
	h2, _ := hexutil.DecodeUint64(last)
	//fmt.Println("last block ", h2)
	return h2
}
