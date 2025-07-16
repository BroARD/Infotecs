package wallet

import (
	"Infotecs/internal/entity"
	"Infotecs/pkg/logging"
	"context"
)

type Service interface {
	Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	GetWalletByID(ctx context.Context, wallet_id string) (entity.Wallet, error)
}


type walletService struct {
	repo Repository
	logger logging.Logger
}

func NewWalletService(repo Repository, logger logging.Logger) Service {
	return &walletService{repo: repo, logger: logger}
}

// Create implements wallet.Service.
func (s *walletService) Create(ctx context.Context, wallet entity.Wallet)  (entity.Wallet, error) {
	return s.repo.Create(ctx, wallet)
}

// GetWalletByID implements wallet.Service.
func (s *walletService) GetWalletByID(ctx context.Context, wallet_id string)  (entity.Wallet, error) {
	return s.repo.GetWalletByID(ctx, wallet_id)
}