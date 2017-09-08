// Copyright 2016 The go-shaft Authors
// This file is part of the go-shaft library.
//
// The go-shaft library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-shaft library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-shaft library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/shaft/go-shaft"

// Verify that Client implements the shaft interfaces.
var (
	_ = shaft.ChainReader(&Client{})
	_ = shaft.TransactionReader(&Client{})
	_ = shaft.ChainStateReader(&Client{})
	_ = shaft.ChainSyncReader(&Client{})
	_ = shaft.ContractCaller(&Client{})
	_ = shaft.GasEstimator(&Client{})
	_ = shaft.GasPricer(&Client{})
	_ = shaft.LogFilterer(&Client{})
	_ = shaft.PendingStateReader(&Client{})
	// _ = shaft.PendingStateEventer(&Client{})
	_ = shaft.PendingContractCaller(&Client{})
)
