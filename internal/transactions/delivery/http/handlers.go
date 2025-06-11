package http

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/transactions"
	"TransactionAPI/internal/transactions/dto"
	"TransactionAPI/pkg/logging"
	"net/http"
	"strconv"

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
	h.logger.Info("Create transaction")
	return func (ctx echo.Context) error {
		getTrans := &dto.CreateTransDTO{}
		ctx.Bind(&getTrans)

		resultTransaction := &models.Transaction{
			ID: uuid.NewString(),
			Status: models.StatusPending,
			Sender: getTrans.From,
			Receiver: getTrans.To,
			Amount: getTrans.Amount,
		}

		createdTrans, err := h.transUC.Create(ctx.Request().Context(), resultTransaction)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		return ctx.JSON(http.StatusCreated, createdTrans)
	}
}

func (h *transHandler) GetByCount() echo.HandlerFunc {
	h.logger.Info("Get transacion by count")
	return func (ctx echo.Context) error{
		trans_count, err := strconv.Atoi(ctx.QueryParam("count"))
		if err != nil {
			return ctx.JSON(http.StatusNotFound, "Param shoud be int")
		}

		transList, err := h.transUC.GetTransactionsByCount(ctx.Request().Context(), trans_count)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, "Could not get transactions list")
		}

		return ctx.JSON(http.StatusOK, transList)
	}
}

