package http

import (
	"TransactionAPI/internal/wallets"

	"github.com/labstack/echo/v4"
)

func MapWalletRoutes(walletGroup *echo.Group, h wallets.Hanlders) {
	walletGroup.GET("/:wallet_id/balance", h.GetByID())
}
