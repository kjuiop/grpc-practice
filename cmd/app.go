package cmd

import (
	"grpc-practice/config"
	"grpc-practice/network"
	"grpc-practice/repository"
	"grpc-practice/service"
	"log"
)

type App struct {
	cfg *config.Config

	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewApp(cfg *config.Config) *App {

	repo, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed init repository, err : %v", err)
	}

	serv, err := service.NewService(cfg, repo)
	if err != nil {
		log.Fatalf("failed init service, err : %v", err)
	}

	net, err := network.NewNetwork(cfg, serv)
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
