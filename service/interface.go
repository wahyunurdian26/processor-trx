package service

import "context"

type TrxProcessorService interface {
	ProcessTransaction(ctx context.Context, trxID string) error
	SubscribeTransaction(ctx context.Context) error
}
