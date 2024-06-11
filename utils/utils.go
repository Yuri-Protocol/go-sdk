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
