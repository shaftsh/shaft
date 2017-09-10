// Copyright 2015 The go-shaft Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main SHAFT network.
var MainnetBootnodes = []string{

	// SHAFT Go Bootnodes
	"enode://42e80282656545c01ab78e47dfa0d664ad5d7468643e81a05b8d17d4c269a9fbf30ffd23cfa065a4599ffa16b7700ed0a0c075c0a027641f9c7e1f765dcad659@158.69.209.240:30307",
	"enode://d4681c54ae0055591f4cf473e85e60ad47c564c7760d2fed7fc25ada06cdb4a845d8f0768e640a0e790cca3e1beb924ec3b4a4e51deb9d169c271bf1b2632f0b@149.56.99.203:30307",
	"enode://155db7f85443ae6b8f68e1dae7c2c89a2dd537239732bf0400ba18324fcc0ee866b2066cbb8b4a1994c13c83cb4e78197041822030a05b257e0144f7f1c61354@137.74.162.146:30307",
	"enode://8401883922960f919b663b2109029cfc3be5cac68a2a6386fdfd404c834fa517c4ca970e8d9463a4b632cfb2031e661cf6630b87f23fff2ee0de86a526a3845f@94.130.105.46:30307",
	"enode://fe4a3fe3822c25052502481954e4ba2088e9662e6059af235e2d7f67ce90a9f30b6dc5478ca2d82fe7e98cc806276c62bf88cd8213bd13c2d1f0199784d6b1b4@94.130.111.226:30307",

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://8401883922960f919b663b2109029cfc3be5cac68a2a6386fdfd404c834fa517c4ca970e8d9463a4b632cfb2031e661cf6630b87f23fff2ee0de86a526a3845f@137.74.162.146:30308",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
