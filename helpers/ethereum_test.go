package helpers

import (
	"reflect"
	"testing"
	
	"github.com/ethereum/go-ethereum/ethclient"
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
			wantErr: false,
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
