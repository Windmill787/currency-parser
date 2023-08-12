package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type BankClient interface {
	ParseRate() (float64, error)
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

var DefaultBankClient = &http.Client{
	Transport: &loggerRoundTripper{
		next:   http.DefaultTransport,
		writer: os.Stdout,
	},
	Timeout: time.Second * 3,
}

func NewClient() *Client {
	return &Client{
		PrivatBankClient: NewPrivatClient(),
	}
}
