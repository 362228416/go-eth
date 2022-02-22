package eth2

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math"
	"math/big"
	"strconv"
)

var (
	client  *ethclient.Client
	_rpc    *rpc.Client
	ChainID int64

	DefaultGasPrice = big.NewInt(30000000000) // 30gwei
)

func InitWeb3(web3Url string, chainId int64) {
	ChainID = chainId
	log.Println("web3 url=", web3Url, "ChainId=", chainId)
	client, _ = ethclient.Dial(web3Url)
	_rpc, _ = rpc.Dial(web3Url)
}

type MetaData struct {
	ABI string
}

func GetTokenBalance(address string, contractAddress string, decimals int) float64 {
	//return 0
	if contractAddress == "" {
		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			return 0
		}
		v := new(big.Float).SetInt(balance)
		f := fmt.Sprintf("%.8f", v.Quo(v, big.NewFloat(math.Pow(10, float64(decimals)))))
		ret, _ := strconv.ParseFloat(f, 64)
		return ret
	}
	token, _ := NewToken20(common.HexToAddress(contractAddress), client)
	balance, _ := token.BalanceOf(nil, common.HexToAddress(address))
	v := new(big.Float).SetInt(balance)
	f := fmt.Sprintf("%.8f", v.Quo(v, big.NewFloat(math.Pow(10, float64(decimals)))))
	ret, _ := strconv.ParseFloat(f, 64)
	return ret
}

func GetTokenBalanceBigInt(address string, contractAddress string) *big.Int {
	if contractAddress == "" {
		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			return big.NewInt(0)
		}
		return balance
	}
	token, _ := NewToken20(common.HexToAddress(contractAddress), client)
	balance, _ := token.BalanceOf(nil, common.HexToAddress(address))
	return balance
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func ToWei(value string, decimals int64) *big.Int {
	amount, _ := Zero().SetString(value)
	val := new(big.Float).Mul(amount, big.NewFloat(math.Pow(10, float64(decimals))))
	val2, _ := new(big.Int).SetString(val.Text('f', 0), 10)
	return val2
}

func FromWei(value *big.Int, decimals int64) float64 {
	v := new(big.Float).SetInt(value)
	w := math.Pow(10, float64(decimals))
	var f = fmt.Sprintf("%.8f", v.Quo(v, big.NewFloat(w)))
	ret, _ := strconv.ParseFloat(f, 64)
	return ret
}

func GetTransactionCount(address string) uint64 {
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return 0
	}
	return nonce
}

func PushTransaction(signTx *types.Transaction) error {
	err := client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Println("send err " + err.Error())
		return err
	}
	return nil
}

func GetWeb3Client() *ethclient.Client {
	return client
}

func GetWeb3RpcClient() *rpc.Client {
	return _rpc
}
