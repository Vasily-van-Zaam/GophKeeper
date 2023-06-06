package main

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/service"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/storage/localstore"
	server "github.com/Vasily-van-Zaam/GophKeeper.git/internal/transport/grpc"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
)

func main() {
	logg := logger.New()
	logg.Info("Server working")
	conf := config.New(logg)
	localstore.New(conf)
	service.New(logg, cryptor.New())
	server.New(logg)

	// conn, err := listener.Accept()
}
