package model

type (
	BaseResponse struct {
		Code    int    `json:"code"  bson:"code"`
		Message string `json:"message"  bson:"message"`
	}
	GetWalletResponse struct {
		Size    int       `json:"size"  bson:"size"`
		Wallets []*Wallet `json:"wallets"  bson:"wallets"`
		BaseResponse
	}
	WalletResponse struct {
		Name        string  `json:"name"  bson:"name"`
		Balance     float64 `json:"balance"  bson:"balance"`
		Coins       []*Coin `json:"coins"  bson:"coins"`
		LastUpdated string  `json:"last_updated"  bson:"last_updated"`
		BaseResponse
	}
	CoinResponse struct {
		Name   string  `json:"name"  bson:"name"`
		Symbol string  `json:"symbol"  bson:"symbol"`
		Amount float64 `json:"amount"  bson:"amount"`
		Rate   float64 `json:"rate"  bson:"rate"`
		BaseResponse
	}
)
