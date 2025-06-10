package http

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/transactions"
	"TransactionAPI/internal/transactions/dto"
	"TransactionAPI/pkg/logging"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type transHandler struct {
	transUC transactions.UseCase
	logger  logging.Logger
}

func NewTransactionHandlers(transUC transactions.UseCase, logger logging.Logger) transactions.Handlers {
	return &transHandler{transUC: transUC, logger: logger}
}

func (h *transHandler) Create() echo.HandlerFunc {
	h.logger.Info("Start creation a transaction")
	return func (c echo.Context) error {
		getTrans := &dto.CreateTransDTO{}
		resultTransaction := models.Transaction{
			ID: uuid.NewString(),
			Sender: getTrans.From,
			Receiver: getTrans.To,
			Amount: getTrans.Amount,
		}

		createdTrans, err := h.transUC.Create(c.Request().Context(), &resultTransaction)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, createdTrans)
	}
}

func (h *transHandler) GetByCount() echo.HandlerFunc {
	return nil
}

