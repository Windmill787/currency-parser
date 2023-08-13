package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "os"
	"time"

	"github.com/Windmill787/currency-parser/entities"
)

type monoCurrency struct {
	BaseCurrency int     `json:"currencyCodeA"`
	Currency     int     `json:"currencyCodeB"`
	Date         int64   `json:"date"`
	SellRate     float64 `json:"rateSell"`
	BuyRate      float64 `json:"rateBuy"`
	CrossRate    float64 `json:"rateCross"`
}

func (m *monoCurrency) Info() string {
	return fmt.Sprintf("[%s] base={%d} curr={%d} sell={%.2f}", time.Unix(m.Date, 0).Format(time.ANSIC), m.BaseCurrency, m.Currency, m.SellRate)
}

type MonoClient struct {
	httpClient *http.Client
}

func NewMonoClient() *MonoClient {
	return &MonoClient{
		httpClient: DefaultBankClient,
	}
}

func (c *MonoClient) ParseRate(currency *entities.Currency) (float64, error) {
	ok := isCurrencyAvailable(currency)
	if !ok {
		return float64(0), fmt.Errorf("currency rate for %s is unavailable", currency.Code)
	}

	resp, err := c.httpClient.Get("https://api.monobank.ua/bank/currency")
	if err != nil {
		return float64(0), err
	}
	defer resp.Body.Close()

	// file, err := os.Open("test/mono.json")
	// if err != nil {
	// 	return float64(0), err
	// }
	// defer file.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return float64(0), err
	}

	var r []monoCurrency
	if err := json.Unmarshal(body, &r); err != nil {
		return float64(0), err
	}

	for _, rate := range r {
		if rate.Currency == baseCurrency.Number && rate.BaseCurrency == currency.Number {
			return rate.SellRate, nil
		}
	}

	return float64(0), fmt.Errorf("currency %s not found", currency.Code)
}
