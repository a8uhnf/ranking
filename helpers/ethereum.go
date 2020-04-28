package helpers

import "github.com/ethereum/go-ethereum/ethclient"

func GetInfuraClient() (*ethclient.Client, error) {
	return ethclient.Dial("wss://ropsten.infura.io/ws/v3/02bad7322763476ea8d4ce4642f3861a")
}

func FindOneMonthAgoBlockNumber()  {
	
}
