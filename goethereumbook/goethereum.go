package goethereumbook

import (
	"context"
	"fmt"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GenerateClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/02bad7322763476ea8d4ce4642f3861a")
	
	return client, err
}

func getHexToAddress(to string) error {
	address := common.HexToAddress(to)
	fmt.Println(address.Hash())
	fmt.Println(string(address.Bytes()))
	return nil
}

func getAccountBalance(to string) error {
	account := common.HexToAddress(to)
	client, err := GenerateClient()
	if err != nil {
		return err
	}
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return err
	}
	fmt.Printf("got balance %v\n", balance)
	return nil
	
}
