package main

import (
	"cryptomon/req"
	"cryptomon/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

const HOUBI_DATA_PATH = "/home/viettr/Desktop/Dev/cryptomon/huobi.json"

func main() {
	tracks := os.Args[1:]
	if len(tracks) == 0 {
		fmt.Println("There is nothing to track")
		os.Exit(0)
	}
	TransactionsFile, err := os.Open(HOUBI_DATA_PATH)
	if err != nil {
		log.Fatalf("Failed to open json file: %s", err)
	}
	defer TransactionsFile.Close()
	transactions := utils.DecodeTransactions(TransactionsFile)

	// Get single crypto price
	coinCap, err := req.GetPrice(tracks[0])
	price, _ := strconv.ParseFloat(coinCap.Data.PriceUSD, 64)
	percentProf := utils.RealTimeProfitLoss(transactions, coinCap.Data.Symbol, price)
	vol, _ := strconv.ParseFloat(coinCap.Data.ChangePer24h, 64)
	fmt.Printf("%-5s %-7.2f %.2f%% %6.2f%%\n", coinCap.Data.Symbol, price, vol, percentProf)

	// Get all crypto prices
	// coinCap, err := req.GetAllPrices()
	// for err != nil || len(coinCap.Data) == 0 {
	// 	os.Exit(1)
	// }
	// CryptoTracks := utils.FindCrypto(tracks, coinCap)
	// for _, crypto := range CryptoTracks {
	// 	price, _ := strconv.ParseFloat(crypto.PriceUSD, 64)
	// 	percentProf := utils.RealTimeProfitLoss(transactions, crypto.Symbol, price)
	// 	vol, _ := strconv.ParseFloat(crypto.ChangePer24h, 64)
	// 	fmt.Printf("%-5s %-10.2f %.2f%% %6.2f%%\n", crypto.Symbol, price, vol, percentProf)
	// }
}
