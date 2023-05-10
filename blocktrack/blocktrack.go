package blocktrack

import (
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockTrack struct {
	CLIENT      *ethclient.Client
	CHAIN_ID    int32
	filterQuery ethereum.FilterQuery
}

// [][]common.Hash{{common.HexToHash("0x0000")}}
func (bt *BlockTrack) Start(chain int32, addresses []common.Address, Topics [][]common.Hash) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/1a2b3c4d5e") // change to websocket
	if err != nil {
		log.Fatal("Failed to connect to websocket RPC ", err)
	} else {
		log.Println("Connected to Ethereum Node")
		bt.CLIENT = client
		bt.CHAIN_ID = chain
		bt.filterQuery = ethereum.FilterQuery{
			Addresses: addresses,
			Topics:    Topics,
		}
	}
}
