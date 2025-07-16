package transaction

import (
	"Infotecs/internal/entity"
	"Infotecs/internal/wallet"
	"Infotecs/pkg/logging"
	"context"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Create(ctx context.Context, trans *entity.Transaction) (*entity.Transaction, error)
	GetTransactionsByCount(ctx context.Context, trans_count int) ([]entity.Transaction, error)
}

type transService struct {
	transRepo Repository
	walletRepo wallet.Repository
	logger    logging.Logger
}

func NewTransService(transRepo Repository, logger logging.Logger, walletRepo wallet.Repository) Service {
	return &transService{transRepo: transRepo, logger: logger, walletRepo: walletRepo}
}

func (s *transService) Create(ctx context.Context, trans *entity.Transaction) (*entity.Transaction, error) {
	//Проверка существует ли отправитель
	sender, errSender := s.walletRepo.GetWalletByID(ctx, trans.Sender)
	if errSender != nil {
		s.logger.Info("Отправителя не существует")
		return nil, echo.NewHTTPError(400, "Отправителя не существует")
	}
	//Проверка существует ли получатель
	receiver, errReceiver := s.walletRepo.GetWalletByID(ctx, trans.Receiver)
	if errReceiver != nil {
		s.logger.Info("Create: Получателя не существует")
		return nil, echo.NewHTTPError(400, "Получателя не существует")
	}
	//Проверка достаточно ли денег на счету
	if sender.Amount < trans.Amount {
		s.logger.Info("Create: Недостаточно денег")
		return nil, echo.NewHTTPError(400, "Недостаточно средств")
	}
	//Списание средств у отправителя
	errAmount := s.walletRepo.UpdateAmount(ctx, sender, sender.Amount-trans.Amount)
	if errAmount != nil {
		trans.Status = entity.StatusFailed
		s.logger.Info("Create: Ошибка при изменении счёта отправителя")
		return s.transRepo.Create(ctx, trans)
	}
	//Зачисление денег получателю
	errAmount = s.walletRepo.UpdateAmount(ctx, receiver, receiver.Amount+trans.Amount)
	if errAmount != nil {
		trans.Status = entity.StatusFailed
		s.logger.Info("Create: Ошибка при изменении счёта получателя")
		return s.transRepo.Create(ctx, trans)
	}
	//Если всё прошло успешно, то меняем статус на complited
	trans.Status = entity.StatusCompleted

	return s.transRepo.Create(ctx, trans)
}

func (s *transService) GetTransactionsByCount(ctx context.Context, trans_count int) ([]entity.Transaction, error) {
	return s.transRepo.GetTransactionsByCount(ctx, trans_count)
}
