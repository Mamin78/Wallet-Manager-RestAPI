package model

type (
	Wallet struct {
		Name        string  `json:"name"  bson:"name"`
		Balance     float64 `json:"balance"  bson:"balance"`
		Coins       []*Coin `json:"coins"  bson:"coins"`
		LastUpdated string  `json:"last_updated"  bson:"last_updated"`
	}
)
