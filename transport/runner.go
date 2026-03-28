package transport

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"microservice/trx-processor/constanta"
	uLog "microservice/util/logger"
)

func (g *GrpcServer) Run() {
	ctx := context.Background()
	uLog.LogInfo(ctx, "Run", "🚀 Server Trx Processor started successfully 🚀")

	if err := g.trxService.SubscribeTransaction(ctx); err != nil {
		uLog.LogError(ctx, "g.trxService.SubscribeTransaction", err)
		return
	}

	uLog.LogInfo(ctx, "Run", constanta.ServiceName+" is running and listening for events")

	go g.waitForShutdown()

	// Wait for stop signal
	<-g.stopChan
	uLog.LogInfo(ctx, "Run", "Server exiting...")
}

func (g *GrpcServer) waitForShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-c
	uLog.LogInfo(context.Background(), "waitForShutdown", "Shutdown signal received")
	g.Stop()
}

func (g *GrpcServer) Stop() {
	ctx := context.Background()
	uLog.LogInfo(ctx, "Stop", "Stopping Server")
	if g.close != nil {
		g.close()
	}
	if g.stopChan != nil {
		close(g.stopChan)
	}
	uLog.LogInfo(ctx, "Stop", "Server stopped successfully")
}
