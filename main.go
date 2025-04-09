package main

import (
	"flag"
	"grpc-practice/cmd"
	"grpc-practice/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()

	a := cmd.NewApp(config.NewConfig(*configFlag))
	a.StartServer()
}
