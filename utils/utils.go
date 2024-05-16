package utils

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const NullAddr = "0x0000000000000000000000000000000000000000"

func GetPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, common.Address{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error casting public key to ECDSA public key")
		return nil, common.Address{}, fmt.Errorf("casting public key to ECDSA public key")
	}
	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
