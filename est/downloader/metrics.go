// Copyright 2015 The go-EvolutionStellarToken Authors
// This file is part of the go-EvolutionStellarToken library.
//
// The go-EvolutionStellarToken library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-EvolutionStellarToken library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-EvolutionStellarToken library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("est/downloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("est/downloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("est/downloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("est/downloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("est/downloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("est/downloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("est/downloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("est/downloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("est/downloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("est/downloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("est/downloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("est/downloader/receipts/timeout", nil)

	stateInMeter   = metrics.NewRegisteredMeter("est/downloader/states/in", nil)
	stateDropMeter = metrics.NewRegisteredMeter("est/downloader/states/drop", nil)

	throttleCounter = metrics.NewRegisteredCounter("est/downloader/throttle", nil)
)
