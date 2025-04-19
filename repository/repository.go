package repository

import (
	"grpc-practice/config"
	"grpc-practice/gRPC/client"
	auth "grpc-practice/gRPC/proto"
)

type Repository struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient
}

func NewRepository(cfg *config.Config, gRPCClient *client.GRPCClient) (*Repository, error) {
	r := &Repository{
		cfg:        cfg,
		gRPCClient: gRPCClient,
	}

	return r, nil
}

func (r *Repository) CreateAuth(name string) (*auth.AuthData, error) {
	return r.gRPCClient.CreateAuth(name)
}
