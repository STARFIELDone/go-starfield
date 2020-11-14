// Copyright 2019 The go-EvolutionStellarToken Authors
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

package est

import (
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/core"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/core/forkid"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/p2p"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/p2p/dnsdisc"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/p2p/enode"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/rlp"
)

// ethEntry is the "est" ENR entry which advertises est protocol
// on the discovery network.
type ethEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e ethEntry) ENRKey() string {
	return "est"
}

// startEthEntryUpdate starts the ENR updater loop.
func (est *EvolutionStellarToken) startEthEntryUpdate(ln *enode.LocalNode) {
	var newHead = make(chan core.ChainHeadEvent, 10)
	sub := est.blockchain.SubscribeChainHeadEvent(newHead)

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case <-newHead:
				ln.Set(est.currentEthEntry())
			case <-sub.Err():
				// Would be nice to sync with est.Stop, but there is no
				// good way to do that.
				return
			}
		}
	}()
}

func (est *EvolutionStellarToken) currentEthEntry() *ethEntry {
	return &ethEntry{ForkID: forkid.NewID(est.blockchain.Config(), est.blockchain.Genesis().Hash(),
		est.blockchain.CurrentHeader().Number.Uint64())}
}

// setupDiscovery creates the node discovery source for the est protocol.
func (est *EvolutionStellarToken) setupDiscovery(cfg *p2p.Config) (enode.Iterator, error) {
	if cfg.NoDiscovery || len(est.config.DiscoveryURLs) == 0 {
		return nil, nil
	}
	client := dnsdisc.NewClient(dnsdisc.Config{})
	return client.NewIterator(est.config.DiscoveryURLs...)
}
