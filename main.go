package main

import (
	"context"
	"fmt"
	"math/big"
	"time"
	
	"github.com/prometheus/common/log"
	
	"github.com/a8uhnf/ranking/helpers"
)

func main() {
	fmt.Println("starting ranking server...")
	
	client, err := helpers.GetInfuraClient()
	if err != nil {
		log.Error(err)
		return
	}
	
	
	var latestBlockNumber *big.Int
	
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("block number: %v\n", block.Number())
	
	latestBlockNumber = block.Number()
	
	timeStart := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -1).Unix()
	// fmt.Printf("latest block number %v\n", latestBlockNumber)
	fmt.Printf("one week ago time: %v\n", oneWeekAgo)
	blockCount := 0
	txCount := 0
	
	for {
		fmt.Printf("new block number: %v\n", latestBlockNumber)
		block, err := client.BlockByNumber(context.Background(), latestBlockNumber)
		if err != nil {
			log.Error(err)
			return
		}
		blockCreateTime := block.Time()
		fmt.Printf("block create time: %v\n", blockCreateTime)
		oneBig := big.NewInt(1)
		fmt.Printf("current unix time: %v\n", time.Now().Unix())
		latestBlockNumber = latestBlockNumber.Sub(latestBlockNumber, oneBig)
		if blockCreateTime < uint64(oneWeekAgo) {
			break
		}
		blockCount++
		txCount += block.Transactions().Len()
	}
	
	timeEnd := time.Now()
	timeTook := timeEnd.Sub(timeStart).Minutes()
	
	fmt.Printf("took %v minutes\n", timeTook)
	fmt.Printf("number of block %v\n", blockCount)
	fmt.Printf("number of transaction %v\n", txCount)
}
