package helpers

import (
	"context"
	"fmt"
	"math/big"
	
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetInfuraClient() (*ethclient.Client, error) {
	return ethclient.Dial("wss://ropsten.infura.io/ws/v3/02bad7322763476ea8d4ce4642f3861a")
}

func GetBlockCreationTime(blockNumber *big.Int, ctx context.Context) (uint64, error) {
	cli, err := GetInfuraClient()
	if err != nil {
		return 0, err
	}
	
	block, err := cli.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return 0, err
	}
	return block.Time(), nil
}

func FindOneMonthAgoBlockNumber(latestBlockNumber *big.Int, oneMonthAgoTime uint64) (*big.Int, error) {
	var high, low *big.Int
	high = latestBlockNumber
	low = big.NewInt(0)
	ctx := context.Background()
	cli, err := GetInfuraClient()
	if err != nil {
		return nil, err
	}
	var mid *big.Int
	mid = big.NewInt(0)
	for {
		if low.Cmp(high) == 1 || low.Cmp(high) == 0 {
			break
		}
		fmt.Printf("high: %v low %v\n", high, low)
		mid.Add(high, low)
		fmt.Printf("mid %v\n", mid)
		mid.Div(mid, big.NewInt(2))
		fmt.Printf("mid %v\n", mid)
		fmt.Printf("high: %v low %v\n", high, low)
		block, err := cli.BlockByNumber(ctx, mid)
		if err != nil {
			return &big.Int{}, err
		}
		
		if block.Time() < uint64(oneMonthAgoTime) {
			low = mid
		} else {
			high = mid
		}
	}
	
	return mid, nil
}
