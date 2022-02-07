package v1

import (
	"github.com/abdullohsattorov/mytaxi-APIGateway/config"
	"github.com/abdullohsattorov/mytaxi-APIGateway/pkg/logger"
	"github.com/abdullohsattorov/mytaxi-APIGateway/services"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
