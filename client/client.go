package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Windmill787/currency-parser/entities"
)

var (
	baseCurrency      = entities.UAH()
	DefaultBankClient = &http.Client{
		Transport: &loggerRoundTripper{
			next:   http.DefaultTransport,
			writer: os.Stdout,
		},
		Timeout: time.Second * 10,
	}
	availableCurrencies = []*entities.Currency{
		entities.USD(),
		entities.EUR(),
	}
)

type BankClient interface {
	ParseRate(currency *entities.Currency) (float64, error)
}

type MonoBankClient interface {
	BankClient
}

type PrivatBankClient interface {
	BankClient
}

type Client struct {
	MonoBankClient
	PrivatBankClient
}

type loggerRoundTripper struct {
	writer io.Writer
	next   http.RoundTripper
}

func (l *loggerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.writer, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func isCurrencyAvailable(currency *entities.Currency) bool {
	for _, c := range availableCurrencies {
		if *c == *currency {
			return true
		}
	}
	return false
}

func NewClient() *Client {
	return &Client{
		PrivatBankClient: NewPrivatClient(),
		MonoBankClient:   NewMonoClient(),
	}
}
