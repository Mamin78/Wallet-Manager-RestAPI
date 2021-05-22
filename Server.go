package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"myapp/handler"
)

func main() {
	e := echo.New()
	h := &handler.Handler{}

	// Route => handler
	e.GET("/wallets", h.GetWallets)
	e.POST("/wallets", h.CreateWallet)
	e.PUT("/wallets/:name", h.EditWallet)
	e.DELETE("/wallets/:name", h.DeleteWallet)

	e.GET("/:walletName", h.GetWallet)
	e.POST("/:walletName/coins", h.CreateCoin)
	e.PUT("/:walletName/:coinName", h.EditCoin)
	e.DELETE("/:walletName/:coinName", h.DeleteCoin)

	e.Logger.Fatal(e.Start(":1373"))
}
