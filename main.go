package main

import (
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Transactions struct {
	txHash    string `gorm:"primaryKey"`
	address   string
	timestamp string
}

func init(db *gorm.DB) {
	// initiate the database connection
	// migrate the database

}

// example indexer function which uses BlockTrack to find all logs related to a tracking address to find out earned tokens
func main() {
	log.Println("Hello World!")
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	bt := BlockTrack{
		CLIENT:   nil,
		CHAIN_ID: 1,
		filterQuery: ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddr},
			Topics:    [][]common.Hash{{common.HexToHash("0x0000")}},
		},
	}

}
