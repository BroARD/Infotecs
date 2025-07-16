package transaction

import (
	"Infotecs/internal/entity"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, trans *entity.Transaction) (*entity.Transaction, error)
	GetTransactionsByCount(ctx context.Context, trans_count int) ([]entity.Transaction, error)
}

type transRepo struct {
	db *gorm.DB
}

func NewTransRepository(db *gorm.DB) Repository {
	return &transRepo{db: db}
}

func (r *transRepo) Create(ctx context.Context, trans *entity.Transaction) (*entity.Transaction, error) {
	err := r.db.WithContext(ctx).Create(trans).Error
	return trans, err
}

func (r *transRepo) GetTransactionsByCount(ctx context.Context, trans_count int) ([]entity.Transaction, error) {
	var transList []entity.Transaction
	err := r.db.WithContext(ctx).Order("created_at DESC").Limit(trans_count).Find(&transList).Error
	return transList, err
}
