package main

import (
	"YellowPepper-FundsTransfers/application"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if err := application.Run(&args); err != nil {
		log.Fatalln(err)
	}
}
