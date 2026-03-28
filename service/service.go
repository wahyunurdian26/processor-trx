package service

import (
	"microservice/trx-processor/repository"
	"microservice/util/broker"
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
