package main

import (
	"YellowPepper-FundsTransfers/application"
	"os"
)

func main() {
	application.Run(os.Args[1:])
}
