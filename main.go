package main

import (
	"cryptomon/req"
	"cryptomon/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tracks := os.Args[1:]
	if len(tracks) == 0 {
		fmt.Println("There is nothing to track")
		os.Exit(0)
	}
	tracks = utils.UpperAll(tracks)
	tmp, _ := req.GetAllPrices()
	CryptoTracks := utils.FindCrypto(tracks, tmp)

	fmt.Printf("%-12s %-12s %s\n", "Currency", "Price", "Change")
	for _, crypto := range CryptoTracks {
		price, _ := strconv.ParseFloat(crypto.PriceUSD, 64)
		vol, _ := strconv.ParseFloat(crypto.ChangePer24h, 64)
		fmt.Printf("%-12s %-12.2f %.2f%%\n", crypto.Symbol, price, vol)
	}
}
