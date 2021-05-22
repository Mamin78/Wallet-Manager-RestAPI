package handler

import (
	"github.com/labstack/echo"
	"myapp/model"
	"net/http"
	"time"
)

func (h *Handler) GetWallet(c echo.Context) (err error) {
	wallet := FindWalletByName(c.Param("walletName"))

	if wallet == nil {
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The wallet not found"
		return c.JSON(http.StatusOK, resp)
	}

	var resp model.WalletResponse
	resp.Name = wallet.Name
	resp.Balance = wallet.Balance
	resp.LastUpdated = wallet.LastUpdated
	resp.Coins = wallet.Coins
	resp.Code = http.StatusOK
	resp.Message = "All coins received successfully!"
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateCoin(c echo.Context) (err error) {
	coin := new(model.Coin)
	if err := c.Bind(&coin); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	walletName := c.Param("walletName")

	if FindWalletByName(c.Param("walletName")) == nil {
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The wallet not found"
		return c.JSON(http.StatusOK, resp)
	}
	if CoinExists(walletName, coin.Name) {
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The Coins exists!"
		return c.JSON(http.StatusOK, resp)
	}

	FindWalletAndAddCoin(walletName, coin)
	UpdateBalance(walletName)
	UpdateLastUpdate(walletName)

	var resp model.CoinResponse
	resp.Name = coin.Name
	resp.Amount = coin.Amount
	resp.Rate = coin.Rate
	resp.Symbol = coin.Symbol
	resp.Code = http.StatusOK
	resp.Message = "Coin added successfully!"
	return c.JSON(http.StatusCreated, resp)
}

func (h *Handler) EditCoin(c echo.Context) (err error) {
	coin := new(model.Coin)
	if err := c.Bind(&coin); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	walletName := c.Param("walletName")
	coinName := c.Param("coinName")
	if !CoinExists(walletName, coinName) {
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The Coins doesn't exists!"
		return c.JSON(http.StatusOK, resp)
	}
	coin = FindCoinAndEditIt(walletName, coinName, coin)
	UpdateBalance(walletName)
	UpdateLastUpdate(walletName)

	var resp model.CoinResponse
	resp.Name = coin.Name
	resp.Amount = coin.Amount
	resp.Rate = coin.Rate
	resp.Symbol = coin.Symbol
	resp.Code = http.StatusOK
	resp.Message = "Coin updated successfully!"
	return c.JSON(http.StatusCreated, resp)
}

func (h *Handler) DeleteCoin(c echo.Context) (err error) {
	walletName := c.Param("walletName")
	coinName := c.Param("coinName")
	if !CoinExists(walletName, coinName) {
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The Coins doesn't exists!"
		return c.JSON(http.StatusOK, resp)
	}

	deleted := FindCoinAndDeleteIt(walletName, coinName)
	UpdateBalance(walletName)
	UpdateLastUpdate(walletName)

	var resp model.CoinResponse
	resp.Name = deleted.Name
	resp.Amount = deleted.Amount
	resp.Rate = deleted.Rate
	resp.Symbol = deleted.Symbol
	resp.Code = http.StatusOK
	resp.Message = "Coin updated successfully!"
	return c.JSON(http.StatusCreated, resp)
}

func FindWalletByName(name string) (wallet *model.Wallet) {
	for i := 0; i < len(wallets); i++ {
		if wallets[i].Name == name {
			return wallets[i]
		}
	}
	return nil
}

func FindWalletByNameIndex(name string) (index int) {
	for i := 0; i < len(wallets); i++ {
		if wallets[i].Name == name {
			return i
		}
	}
	return -1
}

func FindWalletAndAddCoin(name string, coin *model.Coin) (created bool) {
	for i := 0; i < len(wallets); i++ {
		if wallets[i].Name == name {
			wallets[i].Coins = append(wallets[i].Coins, coin)
			return true
		}
	}
	return false
}

func FindCoinAndEditIt(walletName, coinName string, newCoin *model.Coin) (coin *model.Coin) {
	index := FindWalletByNameIndex(walletName)
	for i := 0; i < len(wallets[index].Coins); i++ {
		if wallets[index].Coins[i].Name == coinName {
			wallets[index].Coins[i] = newCoin
			return wallets[index].Coins[i]
		}
	}
	return nil
}

func FindCoinAndDeleteIt(walletName, coinName string) (deleted *model.Coin) {
	index := FindWalletByNameIndex(walletName)
	for i := 0; i < len(wallets[index].Coins); i++ {
		if wallets[index].Coins[i].Name == coinName {
			deleted := wallets[index].Coins[i]
			wallets[index].Coins = append(wallets[index].Coins[:i], wallets[index].Coins[i+1:]...)
			return deleted
		}
	}
	return nil
}

func CoinExists(walletName, coinName string) (exists bool) {
	index := FindWalletByNameIndex(walletName)
	for i := 0; i < len(wallets[index].Coins); i++ {
		if wallets[index].Coins[i].Name == coinName {
			return true
		}
	}
	return false
}

func UpdateBalance(walletName string) {
	index := FindWalletByNameIndex(walletName)
	balance := 0.0
	for i := 0; i < len(wallets[index].Coins); i++ {
		balance += wallets[index].Coins[i].Rate * wallets[index].Coins[i].Amount
	}
	wallets[index].Balance = balance
}
func UpdateLastUpdate(walletName string) {
	index := FindWalletByNameIndex(walletName)
	dt := time.Now()
	wallets[index].LastUpdated = dt.Format("01-02-2006 15:04")
}
