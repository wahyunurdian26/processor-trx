package transport

import (
	"context"
	"os"

	"microservice/trx-processor/config"
	"microservice/trx-processor/repository/micro"
	"microservice/trx-processor/repository/postgres"
	"microservice/trx-processor/service"
	"microservice/util/broker"
	uLog "microservice/util/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcServer struct {
	config     config.Config
	dbPool     *pgxpool.Pool
	brk        broker.MessageBroker
	trxService service.TrxProcessorService
	close      func()
	stopChan   chan struct{}
}

func NewGRPCServer() *GrpcServer {
	ctx := context.Background()
	cfg := config.LoadConfigs()

	dbPool, err := pgxpool.New(ctx, cfg.DatabaseUrl)
	if err != nil {
		uLog.LogError(ctx, "pgxpool.New", err)
		os.Exit(1)
	}

	brk, err := broker.NewRabbitMQBroker(cfg.RabbitMqUrl)
	if err != nil {
		uLog.LogError(ctx, "broker.NewRabbitMQBroker", err)
	}

	connAccount, err := grpc.Dial(cfg.AccountServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		uLog.LogError(ctx, "grpc.Dial AccountService", err)
	}

	repo := postgres.NewTransactionRepository(dbPool)
	accountClient := micro.NewAccountClient(connAccount)
	svc := service.NewTrxProcessorService(repo, accountClient, brk)

	return &GrpcServer{
		config:     cfg,
		dbPool:     dbPool,
		brk:        brk,
		trxService: svc,
		stopChan:   make(chan struct{}),
		close: func() {
			dbPool.Close()
			if connAccount != nil {
				connAccount.Close()
			}
			if brk != nil {
				brk.Close()
			}
		},
	}
}
