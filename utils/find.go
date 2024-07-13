package utils

import (
	"cryptomon/req"
	"strings"
)

type Coin struct {
	ID           string
	Symbol       string
	Name         string
	PriceUSD     string
	ChangePer24h string
}

func FindCrypto(symbols []string, coinCap req.CoinCap) []Coin {
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

func UpperAll(array []string)[]string {
	for i := 0; i < len(array); i++{
		array[i] = strings.ToUpper(array[i])
	}
	return array
}