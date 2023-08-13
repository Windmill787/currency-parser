package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Windmill787/currency-parser/entities"
)

type privatResponse struct {
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

func (c *PrivatClient) ParseRate(currency *entities.Currency) (float64, error) {
	ok := isCurrencyAvailable(currency)
	if !ok {
		return float64(0), fmt.Errorf("currency rate for %s is unavailable", currency.Code)
	}

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?date=%s", "https://api.privatbank.ua/p24api/exchange_rates", time.Now().Format("02.01.2006")))
	if err != nil {
		return float64(0), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return float64(0), err
	}

	var r privatResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return float64(0), err
	}

	for _, rate := range r.ExchangeRate {
		if rate.Currency == currency.Code {
			return rate.SaleRate, nil
		}
	}

	return float64(0), fmt.Errorf("currency %s not found", currency.Code)
}
