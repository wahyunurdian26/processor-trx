package config

import (
	"microservice/util/config"
	"microservice/trx-processor/constanta"
)

type Config struct {
	DatabaseUrl        string
	RabbitMqUrl        string
	AccountServiceAddr string
}

func LoadConfigs() Config {
	return Config{
		DatabaseUrl: config.Get(constanta.DatabaseUrl, "postgres://postgres:postgres@localhost:5432/omnipay_db?sslmode=disable"),
		RabbitMqUrl: config.Get(constanta.RabbitMqUrl, "amqp://guest:guest@localhost:5672/"),
		AccountServiceAddr: config.Get(constanta.AccountServiceAddr, "localhost:6667"),
	}
}
