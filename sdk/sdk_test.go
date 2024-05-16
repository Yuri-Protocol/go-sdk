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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/xcodespace/go-sdk/utils"
)

func NewSdk() (*Erc6551Sdk, error) {
	privateKey, _, err := utils.GetPrivateKey("")
	if err != nil {
		return nil, err
	}
	return NewErc6551Sdk(privateKey, common.HexToAddress("0xA8Be291B81A08166c4A698359b72E6b66a8b054c"),
		common.HexToAddress("0x9F71299ba4322bc449acd75b4Cd8C3776Ef09e24"),
		common.HexToAddress("0x69339d5048dBd6C46A6AAFCb1E4457C4dD678450"),
		"https://data-seed-prebsc-1-s2.bnbchain.org:8545",
		big.NewInt(97))
}
func TestMint721Nft(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	transfer, err := erc6551Sdk.Mint721Nft(big.NewInt(5), "")
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestComputeAccAddr(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	tokenId := big.NewInt(5)
	tblAddr, err := erc6551Sdk.Account(tokenId)
	assert.Nil(t, err)
	t.Logf("tblAddr:%s", tblAddr.String())
}

func TestCreateAccount(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	tokenId := big.NewInt(5)
	transfer, err := erc6551Sdk.CreateAccount(tokenId)
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestTransferNativeToken(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	tblAddr := common.HexToAddress("0x8DA2136FcE19e0FAd3C2364B2c696A2951Ef5d32")
	toAddr := common.HexToAddress("0xd1b9cef657ed8a2159171ddDa2085EFBf4e3A2aB")
	transfer, err := erc6551Sdk.Transfer(common.HexToAddress(utils.NullAddr), tblAddr, toAddr, big.NewInt(700000))
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestTransferErc20Token(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	ercTokenAddr := common.HexToAddress("")
	tblAddr := common.HexToAddress("")
	toAddr := common.HexToAddress("0xd1b9cef657ed8a2159171ddDa2085EFBf4e3A2aB")
	transfer, err := erc6551Sdk.Transfer(ercTokenAddr, tblAddr, toAddr, big.NewInt(60000000))
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestErc6551Sdk_TransferNft(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	toAddr := common.HexToAddress("0xd1b9cef657ed8a2159171ddDa2085EFBf4e3A2aB")
	tokenId := big.NewInt(5)
	transfer, err := erc6551Sdk.TransferNft(toAddr, tokenId)
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}
