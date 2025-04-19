package cmd

import (
	"grpc-practice/config"
	"grpc-practice/gRPC/client"
	"grpc-practice/network"
	"grpc-practice/repository"
	"grpc-practice/service"
	"log"
)

type App struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient
	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewApp(cfg *config.Config) *App {

	grpcClient, err := client.NEWGRPCClient(cfg)
	if err != nil {
		log.Fatalf("failed init grpc client, err : %v", err)
	}

	repo, err := repository.NewRepository(cfg, grpcClient)
	if err != nil {
		log.Fatalf("failed init repository, err : %v", err)
	}

	serv, err := service.NewService(cfg, repo)
	if err != nil {
		log.Fatalf("failed init service, err : %v", err)
	}

	net, err := network.NewNetwork(cfg, serv, grpcClient)
	if err != nil {
		log.Fatalf("failed init network, err : %v", err)
	}

	return &App{
		cfg:        cfg,
		network:    net,
		service:    serv,
		repository: repo,
	}
}

func (a *App) StartServer() {
	a.network.StartServer()
}
