/*
 * Copyright (C) 2024 The Yuri-Protocol Authors
 * This file is part of The Yuri-Protocol library.
 *
 * The Yuri-Protocol is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Yuri-Protocol is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The Yuri-Protocol.  If not, see <http://www.gnu.org/licenses/>.
 */

package sdk

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xcodespace/go-sdk/erc_6551"
	"github.com/xcodespace/go-sdk/utils"
)

type Erc6551Sdk struct {
	privateKey           *ecdsa.PrivateKey
	nft721ContractAddr   common.Address
	registryContractAddr common.Address
	accountContractAddr  common.Address
	address              common.Address
	rpcClient            *ethclient.Client
	chainID              *big.Int
}

func NewErc6551Sdk(privateKey *ecdsa.PrivateKey, nft721ContractAddr, registryContractAddr,
	accountContractAddr common.Address, rpcUrl string, chainID *big.Int) (*Erc6551Sdk, error) {
	rpcClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &Erc6551Sdk{
		privateKey:           privateKey,
		nft721ContractAddr:   nft721ContractAddr,
		registryContractAddr: registryContractAddr,
		accountContractAddr:  accountContractAddr,
		address:              crypto.PubkeyToAddress(privateKey.PublicKey),
		rpcClient:            rpcClient,
		chainID:              chainID,
	}, nil
}
func (this *Erc6551Sdk) Mint721Nft(tokenId *big.Int, uri string) (*types.Transaction, error) {
	erc721Nft, err := erc_6551.NewERC721Asset(this.nft721ContractAddr, this.rpcClient)
	if err != nil {
		return nil, err
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(this.privateKey, this.chainID)
	if err != nil {
		return nil, err
	}
	addrNonce, err := this.rpcClient.PendingNonceAt(context.Background(), this.address)
	if err != nil {
		return nil, err
	}
	transactOpts.From = this.address
	transactOpts.Nonce = big.NewInt(int64(addrNonce))
	transactOpts.Value = big.NewInt(0)
	return erc721Nft.Mint(transactOpts, tokenId, uri)
}

// computes the token bound account address
func (this *Erc6551Sdk) Account(tokenId *big.Int) (common.Address, error) {
	registry, err := erc_6551.NewERC6551Registry(this.registryContractAddr, this.rpcClient)
	if err != nil {
		return common.Address{}, err
	}
	salt := *new([32]byte)
	return registry.Account(nil, this.accountContractAddr, salt, this.chainID, this.nft721ContractAddr, tokenId)
}

func (this *Erc6551Sdk) CreateAccount(tokenId *big.Int) (*types.Transaction, error) {
	registry, err := erc_6551.NewERC6551Registry(this.registryContractAddr, this.rpcClient)
	if err != nil {
		return nil, err
	}
	salt := *new([32]byte)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(this.privateKey, this.chainID)
	if err != nil {
		return nil, err
	}
	addrNonce, err := this.rpcClient.PendingNonceAt(context.Background(), this.address)
	if err != nil {
		return nil, err
	}
	transactOpts.From = this.address
	transactOpts.Nonce = big.NewInt(int64(addrNonce))
	transactOpts.Value = big.NewInt(0)
	return registry.CreateAccount(transactOpts, this.accountContractAddr, salt, this.chainID, this.nft721ContractAddr, tokenId)
}

func (this *Erc6551Sdk) Transfer(tokenAddr, tbaAddr, toAddr common.Address, value *big.Int) (*types.Transaction, error) {
	erc6551Account, err := erc_6551.NewERC6551Account(tbaAddr, this.rpcClient)
	if err != nil {
		return nil, err
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(this.privateKey, this.chainID)
	if err != nil {
		return nil, err
	}
	addrNonce, err := this.rpcClient.PendingNonceAt(context.Background(), this.address)
	if err != nil {
		return nil, err
	}
	transactOpts.From = this.address
	transactOpts.Nonce = big.NewInt(int64(addrNonce))
	transactOpts.Value = big.NewInt(0)
	if tokenAddr == common.HexToAddress(utils.NullAddr) {
		transfer, err := erc6551Account.Execute(transactOpts, toAddr, value, nil, 0)
		if err != nil {
			return nil, err
		}
		return transfer, nil
	} else {
		transfer, err := erc6551Account.TransferToken(transactOpts, tokenAddr, toAddr, value)
		if err != nil {
			return nil, err
		}
		return transfer, nil
	}
}
