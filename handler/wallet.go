package handler

import (
	"github.com/labstack/echo"
	"myapp/model"
	"net/http"
	"time"
)

func (h *Handler) CreateWallet(c echo.Context) (err error) {
	wallet := new(model.Wallet)
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	if WalletExits(wallet.Name){
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The wallet exists"
		return c.JSON(http.StatusOK, resp)
	}
	dt := time.Now()
	wallet.LastUpdated = dt.Format("01-02-2006 15:04")
	wallets = append(wallets, wallet)

	var resp model.WalletResponse
	resp.Name = wallet.Name
	resp.Balance = wallet.Balance
	resp.LastUpdated = wallet.LastUpdated
	resp.Coins = wallet.Coins
	resp.Code = http.StatusOK
	resp.Message = "Food added successfully!"
	return c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetWallets(c echo.Context) (err error) {
	var resp model.GetWalletResponse
	resp.Wallets = wallets
	resp.Size = len(resp.Wallets)
	resp.Code = http.StatusOK
	resp.Message = "All wallets received successfully!"
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) EditWallet(c echo.Context) (err error) {
	wallet := new(model.Wallet)
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	if !WalletExits(c.Param("name")){
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The wallet doesn't exists"
		return c.JSON(http.StatusOK, resp)
	}

	newWallet := ChangeWalletName(c.Param("name"), wallet.Name)

	UpdateLastUpdate(newWallet.Name)
	var resp model.WalletResponse
	resp.Name = newWallet.Name
	resp.Balance = newWallet.Balance
	resp.LastUpdated = newWallet.LastUpdated
	resp.Coins = newWallet.Coins
	resp.Code = http.StatusOK
	resp.Message = "Wallet name changed successfully!"
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteWallet(c echo.Context) (err error) {
	if !WalletExits(c.Param("name")){
		var resp model.BaseResponse
		resp.Code = 400
		resp.Message = "The wallet doesn't exists"
		return c.JSON(http.StatusOK, resp)
	}
	deletedWallet := DeleteWalletByName(c.Param("name"))

	var resp model.WalletResponse
	resp.Name = deletedWallet.Name
	resp.Balance = deletedWallet.Balance
	resp.LastUpdated = deletedWallet.LastUpdated
	resp.Coins = deletedWallet.Coins
	resp.Code = http.StatusOK
	resp.Message = "Wallet deleted (logged out) successfully!"
	return c.JSON(http.StatusOK, resp)
}

func ChangeWalletName(name, newName string) (wallet *model.Wallet) {
	for i := 0; i < len(wallets); i++ {
		if wallets[i].Name == name {
			wallets[i].Name = newName
			return wallets[i]
		}
	}
	return nil
}

func DeleteWalletByName(name string) (deleted *model.Wallet) {
	for i := 0; i < len(wallets); i++ {
		if wallets[i].Name == name {
			temp := wallets[i]
			wallets = append(wallets[:i], wallets[i+1:]...)
			return temp
		}
	}
	return nil
}

func WalletExits(name string) (exists bool) {
	for i := 0; i < len(wallets); i++ {
		if wallets[i].Name == name {
			return true
		}
	}
	return false
}
