package model

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Account struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainId    *big.Int
}

func NewAccount(privateKeyHex string, chainId *big.Int) (*Account, error) {
	ownerPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := ownerPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to cast public key")
	}

	ownerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &Account{
		privateKey: ownerPrivateKey,
		address:    ownerAddress,
		chainId:    chainId,
	}, nil
}

func (a Account) PrivateKey() *ecdsa.PrivateKey {
	return a.privateKey
}

func (a Account) Address() common.Address {
	return a.address
}

func (a Account) ChainId() *big.Int {
	return a.chainId
}
