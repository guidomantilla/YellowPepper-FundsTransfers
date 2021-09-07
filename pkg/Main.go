package main

import (
	"YellowPepper-FundsTransfers/pkg/application"

	"go.uber.org/zap"
)

func main() {
	err := application.Run()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	err = application.Stop()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}
