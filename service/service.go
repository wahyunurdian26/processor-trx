package service

import (
	"github.com/wahyunurdian26/processor-trx/repository"
	"github.com/wahyunurdian26/util/broker"
)

type trxProcessorService struct {
	repo          repository.TransactionRepository
	accountClient repository.AccountClient
	broker        broker.MessageBroker
}

func NewTrxProcessorService(repo repository.TransactionRepository, accountClient repository.AccountClient, b broker.MessageBroker) TrxProcessorService {
	return &trxProcessorService{
		repo:          repo,
		accountClient: accountClient,
		broker:        b,
	}
}
