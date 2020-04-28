package goethereumbook

import "testing"

func Test_generateClient(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "generate client[basic]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateClient(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getHexToAddress(t *testing.T) {
	type args struct {
		to string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "basic hex to address",
			args:    args{to: "0xd0a6053286dc373e72bd24c164bff77587b7241b"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getHexToAddress(tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("getHexToAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
