package service

import (
	"grpc-practice/config"
	auth "grpc-practice/gRPC/proto"
	"grpc-practice/repository"
)

type Service struct {
	cfg *config.Config

	repository *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	s := &Service{
		cfg:        cfg,
		repository: repository,
	}

	return s, nil
}

func (s *Service) CreateAuth(name string) (*auth.AuthData, error) {
	return s.repository.CreateAuth(name)
}
