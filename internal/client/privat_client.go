package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	apiUrl = "https://api.privatbank.ua/p24api/exchange_rates"
)

type privateResponse struct {
	Date            string         `json:"date"`
	Bank            string         `json:"bank"`
	BaseCurrency    int            `json:"baseCurrency"`
	BaseCurrencyLit string         `json:"baseCurrencyLit"`
	ExchangeRate    []exchangeRate `json:"exchangeRate"`
}

type exchangeRate struct {
	BaseCurrency   string  `json:"baseCurrency"`
	Currency       string  `json:"currency"`
	SaleRateNB     float64 `json:"saleRateNB"`
	PurchaseRateNB float64 `json:"purchaseRateNB"`
	SaleRate       float64 `json:"saleRate"`
	PurchaseRate   float64 `json:"purchaseRate"`
}

type PrivatClient struct {
	httpClient *http.Client
}

func NewPrivatClient() *PrivatClient {
	return &PrivatClient{
		httpClient: DefaultBankClient,
	}
}

func (c *PrivatClient) ParseRate(currency string) (float64, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s?date=%s", apiUrl, time.Now().Format("02.01.2006")))
	if err != nil {
		return float64(0), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return float64(0), err
	}

	var pResp privateResponse

	json.Unmarshal(body, &pResp)

	for _, rate := range pResp.ExchangeRate {
		if rate.Currency == currency {
			return rate.SaleRate, nil
		}
	}

	return float64(0), fmt.Errorf("currency %s not found", currency)
}
