package paseto

import (
	"github.com/o1egl/paseto"
	"grpc-practice/config"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

func (p *PasetoMaker) GenerateToken() (string, error) {
	return "", nil
}

func (p *PasetoMaker) VerifyToken(token string) error {
	return nil
}
