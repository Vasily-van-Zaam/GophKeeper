package main

import "github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"

func main() {
	logg := logger.New()
	logg.Info("CLient working")
}
