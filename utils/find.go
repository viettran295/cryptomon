package utils

import (
	"bufio"
	"cryptomon/req"
	"encoding/json"
	"io"
	"log"
	"strings"
)

type Coin struct {
	ID           string
	Symbol       string
	Name         string
	PriceUSD     string
	ChangePer24h string
}

type Transactions struct {
	Pair    string  `json:"Pair"`
	Side    string  `json:"Side"`
	Crypto  float64 `jsong:"Executed"`
	USDT    float64 `json:"Amount"`
	Average float64 `json:"average"`
}

func FindCrypto(symbols []string, coinCap req.AllCoinCap) []Coin {
	var result []Coin
	for _, symbol := range symbols {
		for i := 0; i < len(coinCap.Data); i++ {
			if symbol == coinCap.Data[i].Name || symbol == coinCap.Data[i].Symbol {
				result = append(result, Coin(coinCap.Data[i]))
			}
		}
	}
	return result
}

func UpperAll(array []string) []string {
	for i := 0; i < len(array); i++ {
		array[i] = strings.ToUpper(array[i])
	}
	return array
}

func DecodeTransactions(file io.Reader) []Transactions {
	var trans []Transactions
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tmp Transactions
		line := scanner.Bytes()
		if err := json.Unmarshal(line, &tmp); err != nil {
			log.Println("Error while decoding json data")
			return []Transactions{}
		}
		trans = append(trans, tmp)
	}
	return trans
}

func RealTimeProfitLoss(transactions []Transactions, symbol string, rtPrice float64) float64 {
	for _, trans := range transactions {
		if strings.Contains(trans.Pair, symbol) {
			return ((rtPrice - trans.Average) / trans.Average) * 100
		}
	}
	return 0
}
