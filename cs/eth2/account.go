package eth2

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"strings"
)

type Key struct {
	// to simplify lookups we also store the address
	Address common.Address
	// we only store privkey as pubkey/address can be derived from it
	// privkey in this struct is always in plaintext
	PrivateKey *ecdsa.PrivateKey
}

func NewAddress() (string, string) {
	addr, priv := _getNewAddress()
	for len(priv) != 64 {
		addr, priv = _getNewAddress()
	}
	return addr, priv
}

func _getNewAddress() (string, string) {
	key, _ := newKey(rand.Reader)
	privkey := key.PrivateKey.D.Text(16)
	//fmt.Println("==key.Address.Hash().Hex()", key.Address.Hash().Hex())
	//fmt.Println("==key.Address.Hex()", key.Address.Hex())
	return strings.ToLower(key.Address.Hex()), privkey
}

func newKey(rand io.Reader) (*Key, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand)
	if err != nil {
		return nil, err
	}
	return newKeyFromECDSA(privateKeyECDSA), nil
}

func _hexToECSSA(hexkey string) *ecdsa.PrivateKey {
	k, _ := crypto.HexToECDSA(hexkey)
	return k
}

func _parseKey(hexkey string) *Key {
	k, _ := crypto.HexToECDSA(hexkey)
	key := newKeyFromECDSA(k)
	return key
}

func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *Key {
	key := &Key{
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}

func _parseAddress(hexkey string) (string, string) {
	k, _ := crypto.HexToECDSA(hexkey)
	key := newKeyFromECDSA(k)
	privkey := key.PrivateKey.D.Text(16)
	return strings.ToLower(key.Address.Hex()), privkey
}
