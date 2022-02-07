package main

import (
	"github.com/abdullohsattorov/mytaxi-APIGateway/api"
	"github.com/abdullohsattorov/mytaxi-APIGateway/config"
	"github.com/abdullohsattorov/mytaxi-APIGateway/pkg/logger"
	"github.com/abdullohsattorov/mytaxi-APIGateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
