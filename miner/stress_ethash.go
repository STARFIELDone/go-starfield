// Copyright 2018 The go-EvolutionStellarToken Authors
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

// +build none

// This file contains a miner stress test based on the Ethash consensus engine.
package main

import (
	"crypto/ecdsa"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/accounts/keystore"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/common"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/common/fdlimit"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/consensus/ethash"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/core"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/core/types"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/crypto"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/est"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/est/downloader"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/log"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/miner"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/node"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/p2p"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/p2p/enode"
	"github.com/EvolutionStellarToken/go-EvolutionStellarToken/params"
)

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	fdlimit.Raise(2048)

	// Generate a batch of accounts to seal and fund with
	faucets := make([]*ecdsa.PrivateKey, 128)
	for i := 0; i < len(faucets); i++ {
		faucets[i], _ = crypto.GenerateKey()
	}
	// Pre-generate the ethash mining DAG so we don't race
	ethash.MakeDataset(1, filepath.Join(os.Getenv("HOME"), ".ethash"))

	// Create an Ethash network based off of the Ropsten config
	genesis := makeGenesis(faucets)

	var (
		nodes  []*est.EvolutionStellarToken
		enodes []*enode.Node
	)
	for i := 0; i < 4; i++ {
		// Start the node and wait until it's up
		stack, ethBackend, err := makeMiner(genesis)
		if err != nil {
			panic(err)
		}
		defer stack.Close()

		for stack.Server().NodeInfo().Ports.Listener == 0 {
			time.Sleep(250 * time.Millisecond)
		}
		// Connect the node to all the previous ones
		for _, n := range enodes {
			stack.Server().AddPeer(n)
		}
		// Start tracking the node and its enode
		nodes = append(nodes, ethBackend)
		enodes = append(enodes, stack.Server().Self())

		// Inject the signer key and start sealing with it
		store := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
		if _, err := store.NewAccount(""); err != nil {
			panic(err)
		}
	}

	// Iterate over all the nodes and start mining
	time.Sleep(3 * time.Second)
	for _, node := range nodes {
		if err := node.StartMining(1); err != nil {
			panic(err)
		}
	}
	time.Sleep(3 * time.Second)

	// Start injecting transactions from the faucets like crazy
	nonces := make([]uint64, len(faucets))
	for {
		// Pick a random mining node
		index := rand.Intn(len(faucets))
		backend := nodes[index%len(nodes)]

		// Create a self transaction and inject into the pool
		tx, err := types.SignTx(types.NewTransaction(nonces[index], crypto.PubkeyToAddress(faucets[index].PublicKey), new(big.Int), 21000, big.NewInt(100000000000+rand.Int63n(65536)), nil), types.HomesteadSigner{}, faucets[index])
		if err != nil {
			panic(err)
		}
		if err := backend.TxPool().AddLocal(tx); err != nil {
			panic(err)
		}
		nonces[index]++

		// Wait if we're too saturated
		if pend, _ := backend.TxPool().Stats(); pend > 2048 {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// makeGenesis creates a custom Ethash genesis block based on some pre-defined
// faucet accounts.
func makeGenesis(faucets []*ecdsa.PrivateKey) *core.Genesis {
	genesis := core.DefaultRopstenGenesisBlock()
	genesis.Difficulty = params.MinimumDifficulty
	genesis.GasLimit = 25000000

	genesis.Config.ChainID = big.NewInt(18)
	genesis.Config.EIP150Hash = common.Hash{}

	genesis.Alloc = core.GenesisAlloc{}
	for _, faucet := range faucets {
		genesis.Alloc[crypto.PubkeyToAddress(faucet.PublicKey)] = core.GenesisAccount{
			Balance: new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
		}
	}
	return genesis
}

func makeMiner(genesis *core.Genesis) (*node.Node, *est.EvolutionStellarToken, error) {
	// Define the basic configurations for the EvolutionStellarToken node
	datadir, _ := ioutil.TempDir("", "")

	config := &node.Config{
		Name:    "geth",
		Version: params.Version,
		DataDir: datadir,
		P2P: p2p.Config{
			ListenAddr:  "0.0.0.0:0",
			NoDiscovery: true,
			MaxPeers:    25,
		},
		NoUSB:             true,
		UseLightweightKDF: true,
	}
	// Create the node and configure a full EvolutionStellarToken node on it
	stack, err := node.New(config)
	if err != nil {
		return nil, nil, err
	}
	ethBackend, err := est.New(stack, &est.Config{
		Genesis:         genesis,
		NetworkId:       genesis.Config.ChainID.Uint64(),
		SyncMode:        downloader.FullSync,
		DatabaseCache:   256,
		DatabaseHandles: 256,
		TxPool:          core.DefaultTxPoolConfig,
		GPO:             est.DefaultConfig.GPO,
		Ethash:          est.DefaultConfig.Ethash,
		Miner: miner.Config{
			GasFloor: genesis.GasLimit * 9 / 10,
			GasCeil:  genesis.GasLimit * 11 / 10,
			GasPrice: big.NewInt(1),
			Recommit: time.Second,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	err = stack.Start()
	return stack, ethBackend, err
}
