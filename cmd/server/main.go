package main

import (
	"log"
	"net"

	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
)

func main() {
	logg := logger.New()
	logg.Info("Server working")

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// conn, err := listener.Accept()
}
