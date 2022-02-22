package eth2

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

func SendToAddress(to string, amount *big.Int, _nonce int, hexkey string) string {
	signTx := CreateSignTransaction(to, amount, _nonce, hexkey)
	//fmt.Println(common.Bytes2Hex(signTx.Data()))
	bytes, _ := signTx.MarshalJSON()
	fmt.Println(string(bytes))
	err := client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Println("send err ", err.Error())
		for i := 0; i < 30; i++ {
			err = client.SendTransaction(context.Background(), signTx)
			if err == nil {
				break
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}
	return signTx.Hash().Hex()
}

func CreateSignTransaction(to string, amount *big.Int, _nonce int, hexkey string) *types.Transaction {
	nonce := uint64(_nonce)
	gas := uint64(21000)
	gasPrice := DefaultGasPrice
	data := []byte{}

	tx := types.NewTransaction(nonce, common.HexToAddress(to), amount, gas, gasPrice, data)
	signTx := SignTx(tx, hexkey)
	return signTx
}

func SignTx(tx *types.Transaction, hexkey string) *types.Transaction {
	key := _parseKey(hexkey)
	signer := types.NewEIP155Signer(big.NewInt(ChainID))
	signedTx, err := types.SignTx(tx, signer, key.PrivateKey)
	if err != nil {
		log.Println("sign error == ", err.Error())
		return nil
	}
	return signedTx
}
