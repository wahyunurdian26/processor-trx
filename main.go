package main

import (
	"github.com/wahyunurdian26/processor-trx/transport"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	srv := transport.NewGRPCServer()
	srv.Run()
}
