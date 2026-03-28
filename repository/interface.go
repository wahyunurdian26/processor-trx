package repository

import (
	"context"
	"microservice/trx-processor/model"
)

type TransactionRepository interface {
	UpdateStatus(ctx context.Context, id string, status string) error
	FindByID(ctx context.Context, id string) (*model.Transaction, error)
}

type AccountClient interface {
	DeductBudget(ctx context.Context, accountID string, amount float64) error
}
