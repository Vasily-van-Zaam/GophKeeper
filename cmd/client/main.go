package main

import (
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
)

func main() {
	versionClient := "0.0.1"
	token := "secret_key_version_0.0.1"
	conf := config.New(logger.New(), cryptor.New(), versionClient, token)

	client, err := appclient.New(conf)

	if err != nil {
		conf.Logger().Fatal(err, client)
		return
	}
	log.Fatal(client.Run())
}
