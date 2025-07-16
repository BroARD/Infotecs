package wallet

import (
	"Infotecs/internal/entity"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	GetWalletByID(ctx context.Context, wallet_id string) (entity.Wallet, error)
	UpdateAmount(ctx context.Context, wallet entity.Wallet, new_amount float64) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) Repository {
	return &walletRepository{db: db}
}

func (r *walletRepository) Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	err := r.db.WithContext(ctx).Create(wallet).Error
	return wallet, err
}

func (r *walletRepository) GetWalletByID(ctx context.Context, wallet_id string) (entity.Wallet, error) {
	var wallet entity.Wallet
	err := r.db.WithContext(ctx).First(&wallet, "id = ?", wallet_id).Error
	return wallet, err
}

func (r *walletRepository) UpdateAmount(ctx context.Context, wallet entity.Wallet, new_amount float64) error {
	err := r.db.WithContext(ctx).Model(wallet).Update("amount", new_amount).Error
	return err
}
