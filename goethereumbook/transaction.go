package goethereumbook

import "context"

func GetLatestBlockHeader() (string, error) {
	client, err := GenerateClient()
	if err != nil {
		return "", err
	}
	block, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", err
	}
	return block.Number.String(), nil
}
