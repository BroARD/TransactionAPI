package usecase

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/transactions"
	"TransactionAPI/internal/wallets"
	"TransactionAPI/pkg/logging"
	"context"

	"github.com/labstack/echo/v4"
)

type transUC struct {
	transRepo transactions.Repository
	walletRepo wallets.Repository
	logger    logging.Logger
}

func NewTransUseCase(transRepo transactions.Repository, logger logging.Logger, walletRepo wallets.Repository) transactions.UseCase {
	return &transUC{transRepo: transRepo, logger: logger, walletRepo: walletRepo}
}

func (u *transUC) Create(ctx context.Context, trans *models.Transaction) (*models.Transaction, error) {
	//Проверка существует ли отправитель
	sender, errSender := u.walletRepo.GetWalletByID(ctx, trans.Sender)
	if errSender != nil {
		return nil, errSender
	}
	//Проверка существует ли получатель
	receiver, errReceiver := u.walletRepo.GetWalletByID(ctx, trans.Receiver)
	if errReceiver != nil {
		return nil, errReceiver
	}
	//Проверка достаточно ли денег на счету
	if sender.Amount < trans.Amount {
		return nil, echo.NewHTTPError(400, "Insufficient funds")
	}
	//Списание средств у отправителя
	errAmount := u.walletRepo.UpdateAmount(ctx, sender, sender.Amount-trans.Amount)
	if errAmount != nil {
		trans.Status = models.StatusFailed
		return u.transRepo.Create(ctx, trans)
	}
	//Зачисление денег получателю
	errAmount = u.walletRepo.UpdateAmount(ctx, receiver, receiver.Amount+trans.Amount)
	if errAmount != nil {
		trans.Status = models.StatusFailed
		return u.transRepo.Create(ctx, trans)
	}
	//Если всё прошло успешно, то меняем статус на complited
	trans.Status = models.StatusCompleted

	return u.transRepo.Create(ctx, trans)
}

func (u *transUC) GetTransactionsByCount(ctx context.Context, trans_count int) ([]models.Transaction, error) {
	return u.transRepo.GetTransactionsByCount(ctx, trans_count)
}
