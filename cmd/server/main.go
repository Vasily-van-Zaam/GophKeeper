package main

import (
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/service"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/storage/pstgql"
	server "github.com/Vasily-van-Zaam/GophKeeper.git/internal/transport/grpc"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
)

func main() {
	logg := logger.New()
	logg.Info("Server working")
	conf := config.New(logg, cryptor.New())
	store, err := pstgql.New(conf)
	if err != nil {
		log.Fatal(err)
	}
	managerService := service.New(conf, store, cryptor.New())
	userSrvice := service.NewUserService(conf, store, cryptor.New())
	srv := server.New(conf, userSrvice, managerService)

	log.Fatal(srv.Run(conf.Server().RunAddrss()))
}
