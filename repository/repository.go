package repository

import "grpc-practice/config"

type Repository struct {
	cfg *config.Config
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{
		cfg: cfg,
	}

	return r, nil
}
