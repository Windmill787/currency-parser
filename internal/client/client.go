package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type loggerRoundTripper struct {
	writer io.Writer
	next   http.RoundTripper
}

func (l *loggerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.writer, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func NetClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect")
			return nil
		},
		Transport: &loggerRoundTripper{
			next:   http.DefaultTransport,
			writer: os.Stdout,
		},
		Timeout: time.Second * 3,
	}
}
