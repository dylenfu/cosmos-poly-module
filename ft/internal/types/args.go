/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"fmt"
	"github.com/polynetwork/cosmos-poly-module/common"
	polycommon "github.com/polynetwork/poly/common"
	"math/big"
)

type TxArgs struct {
	ToAddress []byte
	Amount    *big.Int
}

func (this *TxArgs) Serialization(sink *polycommon.ZeroCopySink, intBsLen int) error {
	sink.WriteVarBytes(this.ToAddress)
	paddedAmountBs, err := common.PadFixedBytes(this.Amount, intBsLen)
	if err != nil {
		return fmt.Errorf("TxArgs Serialization error:%v", err)
	}
	sink.WriteBytes(polycommon.ToArrayReverse(paddedAmountBs))
	return nil
}

func (this *TxArgs) Deserialization(source *polycommon.ZeroCopySource, intBsLen int) error {
	toAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("TxArgs deserialize ToAddress error")
	}
	paddedAmountBs, eof := source.NextBytes(uint64(intBsLen))
	if eof {
		return fmt.Errorf("TxArgs deserialize Amount error")
	}
	amount, err := common.UnpadFixedBytes(paddedAmountBs, intBsLen)
	if err != nil {
		return fmt.Errorf("TxArgs Deserialization error:%v", err)
	}

	this.ToAddress = toAddress
	this.Amount = amount
	return nil
}
