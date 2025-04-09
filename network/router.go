package network

import (
	"github.com/gin-gonic/gin"
	"grpc-practice/config"
	"grpc-practice/service"
	"log"
)

type Network struct {
	cfg *config.Config

	service *service.Service
	engine  *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	n := &Network{
		cfg:     cfg,
		service: service,
		engine:  gin.New(),
	}

	return n, nil
}

func (n *Network) StartServer() {
	if err := n.engine.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
