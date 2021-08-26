package main

import (
	"YellowPepper-FundsTransfers/application"
	"os"
)

func main() {
	args := os.Args[1:]
	application.Run(&args)
}
