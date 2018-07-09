package main

import (
	"fmt"

	"./turtlecoind"
	"./walletd"
)

func main() {
	walletdResponse := walletd.GetStatus("passw0rd", "localhost", 8070)
	turtlecoindResponse := turtlecoind.GetBlockCount("localhost", 11898)
	fmt.Println(walletdResponse)
	fmt.Println(turtlecoindResponse)
}
