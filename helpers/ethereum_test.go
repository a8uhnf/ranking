package helpers

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"testing"
	"time"
	
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestGetInfuraClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *ethclient.Client
		wantErr bool
	}{
		{
			name:    "generate client",
			want:    &ethclient.Client{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInfuraClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfuraClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfuraClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBlockCreationTime(t *testing.T) {
	type args struct {
		blockNumber *big.Int
		ctx         context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "get block creation time 1 [basic]",
			args: args{
				blockNumber: big.NewInt(7799067),
				ctx:         context.Background(),
			},
			want:    1588001834,
			wantErr: false,
		},
		{
			name: "get block creation time 2 [basic]",
			args: args{
				blockNumber: big.NewInt(7799068),
				ctx:         context.Background(),
			},
			want:    1588001885,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBlockCreationTime(tt.args.blockNumber, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockCreationTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBlockCreationTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOneMonthAgoBlockNumber(t *testing.T) {
	type args struct {
		latestBlockNumber        *big.Int
		currentBlockCreationTime uint64
	}
	
	var ar []args
	ctx := context.Background()
	cli, err := GetInfuraClient()
	require.Nil(t, err, "get infura client. error should be nil")
	
	for i := 0; i < 3; i++ {
		block, err := cli.BlockByNumber(ctx, nil)
		require.Nil(t, err, "latest block error should be nil")
		ar = append(ar, args{
			latestBlockNumber:        block.Number(),
			currentBlockCreationTime: uint64(time.Now().AddDate(0, -1, 0).Unix()),
		})
	}
	require.NotZero(t, len(ar), "argument array should not be zero")
	require.Equal(t, 3, len(ar), "length should be three")
	
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name:    "get one month ago block number 1",
			args:    ar[0],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindOneMonthAgoBlockNumber(tt.args.latestBlockNumber, tt.args.currentBlockCreationTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOneMonthAgoBlockNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.NotNil(t, got, "one month ago block number shouldn't be nil...")
			// main test
			one := big.NewInt(1)
			oneLess := big.NewInt(0)
			oneLess.Sub(got, one)
			oneMore := big.NewInt(0)
			oneMore.Add(got, one)
			fmt.Printf("got %v one-more %v one-less %v\n", got, oneMore, oneLess)
			b1, err := cli.BlockByNumber(ctx, oneLess)
			require.Nil(t, err, "getting one less block, shouldn't be nil")
			require.LessOrEqual(t, tt.args.currentBlockCreationTime, b1.Time(), "one less block number should be less than current block creation time")
			b2, err := cli.BlockByNumber(ctx, oneMore)
			require.Nil(t, err, "one more block: error should be nil")
			
			require.GreaterOrEqual(t, b2.Time(), tt.args.currentBlockCreationTime, "one more block number should be less than current block creation time")
		})
	}
}
