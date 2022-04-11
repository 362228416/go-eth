package eth2

import (
	"bytes"
	"context"
	"encoding/hex"
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

func CreateTokenSignTx(priv string, to string, contractAddress string, value *big.Int, nonce int) *types.Transaction {
	// 手动构建input
	methodId := common.FromHex("0xa9059cbb")
	var data []byte
	data = append(data, methodId...)
	paddedAddress := common.LeftPadBytes(common.HexToAddress(to).Bytes(), 32)
	data = append(data, paddedAddress...)
	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
	data = append(data, paddedAmount...)

	gas := uint64(65000)
	//gasPrice := big.NewInt(1000000000)
	//gasPrice, _ := client.SuggestGasPrice(context.Background())
	var gasPrice *big.Int
	for i := 0; i < 10; i++ {
		gasPrice, _ = client.SuggestGasPrice(context.Background())
		if gasPrice != nil && gasPrice.Cmp(big.NewInt(0)) > 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
	if gasPrice == nil || gasPrice.Cmp(big.NewInt(0)) <= 0 {
		log.Println("获取不到GasPrice，默认5gwei")
		gasPrice = big.NewInt(5000000000)
	}
	tx := types.NewTransaction(uint64(nonce), common.HexToAddress(contractAddress), big.NewInt(0), gas, gasPrice, data)
	signTx := SignTx(tx, priv)

	log.Println("create token tx chainId=", ChainID, "nonce=", nonce, "gasPrice=", gasPrice, "value=", value, "to=", to)

	w := new(bytes.Buffer)
	signTx.EncodeRLP(w)
	rawTx := hex.EncodeToString(w.Bytes())
	log.Println("rawTx=", "0x"+rawTx)

	return signTx

}
