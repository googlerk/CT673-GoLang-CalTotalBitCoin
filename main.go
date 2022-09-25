package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var minedSatoshi = 50*10 ^ 8
	var numBlockIn4Years = 6 * 24 * 365 * 4
	var sumBitcoin = 0
	for {
		if minedSatoshi <= 0 {
			break
		}
		sumBitcoin += minedSatoshi * numBlockIn4Years
		minedSatoshi /= 2
	}

	fmt.Println("")
	fmt.Printf("Total BitCoin in satoshi : %v \n", sumBitcoin)
	fmt.Printf("Total BitCoin : %v \n", sumBitcoin/10^8)
	fmt.Println("")
	fmt.Println("Press any key to exit.")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		return
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
		return
	}
}
