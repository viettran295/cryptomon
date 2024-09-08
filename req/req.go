package req

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	URLAllCoins string = "https://api.coincap.io/v2/assets/"
	URLCoin     string = "https://api.coincap.io/v2/assets/"
)

type CoinCapData struct {
	ID           string `json:"id"`
	Symbol       string `json:"Symbol"`
	Name         string `json:"name"`
	PriceUSD     string `json:"priceUsd"`
	ChangePer24h string `json:"changePercent24Hr"`
}
type AllCoinCap struct {
	Data []CoinCapData `json:"data"`
}
type SingleCoinCap struct {
	Data CoinCapData `json:"data"`
}

func httpRequest(url string, symbol string) (*http.Response, error) {
	apiKey := os.Getenv("COINCAP_KEY")
	bearerWKey := "Bearer " + apiKey
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url+symbol, nil)
	req.Header.Set("Authorization", bearerWKey)
	return client.Do(req)
}

func GetAllPrices() (AllCoinCap, error) {
	resp, err := httpRequest(URLAllCoins, "")
	if err != nil || resp.Body == nil || resp.StatusCode != 200 {
		log.Println(resp.Status)
		return AllCoinCap{}, err
	}
	defer resp.Body.Close()

	var data AllCoinCap
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error while processing json data")
		return AllCoinCap{}, err
	}
	return data, nil
}

func GetPrice(symbol string) (SingleCoinCap, error) {
	resp, err := httpRequest(URLCoin, symbol)
	if err != nil || resp.Body == nil || resp.StatusCode != 200 {
		log.Println(resp.Status)
		return SingleCoinCap{}, err
	}
	defer resp.Body.Close()

	var data SingleCoinCap
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error while processing json data")
		return SingleCoinCap{}, err
	}
	return data, nil
}
