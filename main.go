package main

import (
	"flag"
	"grpc-practice/cmd"
	"grpc-practice/config"
	"grpc-practice/gRPC/server"
	"log"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*configFlag)

	a := cmd.NewApp(cfg)

	if err := server.NewGRPCServer(cfg); err != nil {
		log.Fatalf("failed new grpc server, err : %w", err)
	}
	time.Sleep(1e9)

	a.StartServer()
}
