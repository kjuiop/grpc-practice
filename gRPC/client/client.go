package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-practice/config"
	"grpc-practice/gRPC/paseto"
	auth "grpc-practice/gRPC/proto"
	"time"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NEWGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	if client, err := grpc.NewClient(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.authClient = auth.NewAuthServiceClient(client)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}

	return c, nil
}

func (g *GRPCClient) CreateAuth(name string) (*auth.AuthData, error) {
	// CreateDate 를 넣어주기 위함
	now := time.Now()
	expiredTime := now.Add(30 * time.Minute)

	a := &auth.AuthData{
		Name:       name,
		CreateDate: now.Unix(),
		ExpireDate: expiredTime.Unix(),
	}

	if token, err := g.pasetoMaker.GenerateToken(a); err != nil {
		return nil, err
	} else {
		a.Token = token

		if res, err := g.authClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: a}); err != nil {
			return nil, err
		} else {
			return res.Auth, nil
		}
	}
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.Verify, error) {

	if res, err := g.authClient.VerifyAuth(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return res.V, nil
	}
}
