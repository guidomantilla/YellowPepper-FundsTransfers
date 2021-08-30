package main

import (
	"YellowPepper-FundsTransfers/pkg/application"
	"log"
)

func main() {
	if err := application.Run(); err != nil {
		log.Fatalln(err)
	}
}
