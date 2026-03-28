package postgres

import (
	"context"
	"errors"
	"time"

	"microservice/trx-processor/model"
	"microservice/trx-processor/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepository struct {
	db *pgxpool.Pool
}

const timeout = 5

func NewTransactionRepository(db *pgxpool.Pool) repository.TransactionRepository {
	return &transactionRepository{db: db}
}

var (
	updateStatusQuery = `UPDATE transactions SET status = $1 WHERE id = $2`
	findByIDQuery     = `SELECT id, account_id, amount, merchant_name, description, status, created_at FROM transactions WHERE id = $1`
)

func (r *transactionRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()
	_, err := r.db.Exec(ctxTimeout, updateStatusQuery, status, id)
	return err
}

func (r *transactionRepository) FindByID(ctx context.Context, id string) (*model.Transaction, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()
	row := r.db.QueryRow(ctxTimeout, findByIDQuery, id)
	var trx model.Transaction
	err := row.Scan(&trx.ID, &trx.AccountID, &trx.Amount, &trx.MerchantName, &trx.Description, &trx.Status, &trx.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &trx, nil
}
