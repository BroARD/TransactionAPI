package http

import (
	"TransactionAPI/internal/wallets"
	"TransactionAPI/pkg/logging"
	"net/http"

	"github.com/labstack/echo/v4"
)

type walletHandlers struct {
	walletUC wallets.UseCase
	logger   logging.Logger
}


func NewWalletHandlers(walletUC wallets.UseCase, logger logging.Logger) wallets.Hanlders {
	return &walletHandlers{walletUC: walletUC, logger: logger}
}

// Create implements wallets.Hanlders.
func (h *walletHandlers) GetByID() echo.HandlerFunc {
	h.logger.Info("Get transacion by count")
	return func(ctx echo.Context) error {
		wallet_id := ctx.Param("wallet_id")

		wallet, err := h.walletUC.GetWalletByID(ctx.Request().Context(), wallet_id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, "This id is not exists")
		}
		return ctx.JSON(http.StatusOK, wallet)
	}

}
