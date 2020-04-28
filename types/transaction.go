package types

type Transaction struct {
	Nonce    string `json:"nonce"`
	GasPrice string `json:"gasPrice"`
	Gas      string `json:"gas"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Input    string `json:"input"`
	V        string `json:"v"`
	R        string `json:"r"`
	S        string `json:"s"`
	Hash     string `json:"hash"`
}
