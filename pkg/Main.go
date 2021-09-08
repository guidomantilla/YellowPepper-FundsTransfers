package main

import (
	"YellowPepper-FundsTransfers/pkg/application"
	"log"
)

func main() {

	if err := application.Run(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := application.Stop(); err != nil {
		log.Fatalln(err.Error())
	}
}
