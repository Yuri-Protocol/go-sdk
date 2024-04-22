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
	return NewErc6551Sdk(nil, common.HexToAddress(""),
		common.HexToAddress(""),
		common.HexToAddress(""),
		"https://bsc-testnet-rpc.publicnode.com",
		big.NewInt(97))
}
func TestMint721Nft(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	transfer, err := erc6551Sdk.Mint721Nft(big.NewInt(1), "")
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestComputeAccAddr(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	tblAddr, err := erc6551Sdk.Account(big.NewInt(1))
	assert.Nil(t, err)
	t.Logf("tblAddr:%s", tblAddr.String())
}

func TestCreateAccount(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	transfer, err := erc6551Sdk.CreateAccount(big.NewInt(1))
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestTransferNativeToken(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	tblAddr := common.HexToAddress("")
	toAddr := common.HexToAddress("")
	transfer, err := erc6551Sdk.Transfer(common.HexToAddress(utils.NullAddr), tblAddr, toAddr, big.NewInt(1000))
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}

func TestTransferErc20Token(t *testing.T) {
	erc6551Sdk, err := NewSdk()
	assert.Nil(t, err)
	ercTokenAddr := common.HexToAddress("")
	tblAddr := common.HexToAddress("")
	toAddr := common.HexToAddress("")
	transfer, err := erc6551Sdk.Transfer(ercTokenAddr, tblAddr, toAddr, big.NewInt(1000))
	assert.Nil(t, err)
	t.Logf("txHash:%s", transfer.Hash().String())
}
