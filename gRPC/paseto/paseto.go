package paseto

import (
	"crypto/rand"
	"github.com/o1egl/paseto"
	"grpc-practice/config"
	auth "grpc-practice/gRPC/proto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

func (p *PasetoMaker) GenerateToken(auth *auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return p.Pt.Encrypt(p.Key, auth, randomBytes)
}

func (p *PasetoMaker) VerifyToken(token string) error {
	var a *auth.AuthData
	return p.Pt.Decrypt(token, p.Key, a, nil)
}
