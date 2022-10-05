package models

type Eth struct {
	Ethereum Transactions `json:"ethereum"`
}

type Transactions struct {
	Transactions []Info `json:"transactions"`
}

type Info struct {
	Time           string  `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}
