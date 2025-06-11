package http

import (
	"TransactionAPI/internal/transactions"

	"github.com/labstack/echo/v4"

)

func MapTransationsRoutes(transGroup *echo.Group, h transactions.Handlers) {
	transGroup.POST("/send", h.Create())
	transGroup.GET("/transactions", h.GetByCount())
}