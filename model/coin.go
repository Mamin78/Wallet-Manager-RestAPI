package model

type (
	Coin struct {
		Name   string  `json:"name"  bson:"name"`
		Symbol string  `json:"symbol"  bson:"symbol"`
		Amount float64 `json:"amount"  bson:"amount"`
		Rate   float64 `json:"rate"  bson:"rate"`
	}
)
