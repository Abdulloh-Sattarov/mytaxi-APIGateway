package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/abdullohsattorov/mytaxi-APIGateway/config"
	pb "github.com/abdullohsattorov/mytaxi-APIGateway/genproto"
)

type IServiceManager interface {
	MyTaxiService() pb.MyTaxiServiceClient
}

type serviceManager struct {
	taxiService pb.MyTaxiServiceClient
}

func (s *serviceManager) MyTaxiService() pb.MyTaxiServiceClient {
	return s.taxiService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connMyTaxi, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.MyTaxiServiceHost, conf.MyTaxiServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taxiService: pb.NewMyTaxiServiceClient(connMyTaxi),
	}

	return serviceManager, nil
}
