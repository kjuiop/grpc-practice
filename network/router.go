package network

import (
	"github.com/gin-gonic/gin"
	"grpc-practice/config"
	"grpc-practice/gRPC/client"
	"grpc-practice/service"
	"log"
)

type Network struct {
	cfg *config.Config

	service    *service.Service
	gRPCClient *client.GRPCClient
	engine     *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service, grpcClient *client.GRPCClient) (*Network, error) {
	n := &Network{
		cfg:        cfg,
		service:    service,
		gRPCClient: grpcClient,
		engine:     gin.New(),
	}

	// 1. token 생성 api
	n.engine.POST("/login", n.Login)

	n.engine.GET("/verify", n.verifyLogin(), n.Verify)

	// 2. token 검증 api

	return n, nil
}

func (n *Network) StartServer() {
	if err := n.engine.Run(":8081"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
