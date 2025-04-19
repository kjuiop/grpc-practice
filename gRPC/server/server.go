package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-practice/config"
	"grpc-practice/gRPC/paseto"
	auth "grpc-practice/gRPC/proto"
	"net"
	"time"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {

		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})

		reflection.Register(grpc.NewServer())

		go func() {
			if err := server.Serve(lis); err != nil {
				panic(err)
			}
		}()

	}

	return nil
}

func (g *GRPCServer) CreateAuth(ctx context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token
	g.tokenVerifyMap[token] = data

	return &auth.CreateTokenRes{Auth: data}, nil
}

func (g *GRPCServer) VerifyAuth(ctx context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token
	res := &auth.VerifyTokenRes{
		V: &auth.Verify{
			Auth: nil,
		},
	}

	if authData, ok := g.tokenVerifyMap[token]; !ok {
		res.V.Status = auth.ResponseType_FAILED
		return res, errors.New("not Existed At TokenVerifyMap")
	} else if err := g.pasetoMaker.VerifyToken(token); err != nil {
		return res, fmt.Errorf("failed verify token, err : %w", err)
	} else if authData.ExpireDate < time.Now().Unix() {
		// 만료된 토큰이기 때문에 재 로그인을 위해 토큰 삭제
		delete(g.tokenVerifyMap, token)
		res.V.Status = auth.ResponseType_EXPIRED_DATE
		return res, errors.New("expired Token")
	} else {
		res.V.Status = auth.ResponseType_SUCCESS
		return res, nil
	}
}

func (g *GRPCServer) mustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
