package main

import (
	"microservice/trx-processor/transport"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	srv := transport.NewGRPCServer()
	srv.Run()
}
