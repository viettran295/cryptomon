package req

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const URL string = "https://api.coincap.io/v2/assets"

type CoinCap struct {
	Data []struct {
		ID           string `json:"id"`
		Symbol       string `json:"Symbol"`
		Name         string `json:"name"`
		PriceUSD     string `json:"priceUsd"`
		ChangePer24h string `json:"changePercent24Hr"`
	}
}

func GetAllPrices() (CoinCap, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can not loat env variables")
	}
	apiKey := os.Getenv("COINCAP_KEY")
	bearerWKey := "Bearer " + apiKey
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", bearerWKey)
	resp, err := client.Do(req)
	if err != nil || resp.Body == nil || resp.StatusCode != 200 {
		log.Println(resp.Status)
		return CoinCap{}, err
	}
	defer resp.Body.Close()

	var data CoinCap
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error while processing json data")
		return CoinCap{}, err
	}
	return data, nil

}
