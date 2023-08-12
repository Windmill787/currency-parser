package service

import (
	"github.com/Windmill787/currency-parser/internal/client"
)

type Service struct {
	client *client.Client
}

func NewService(client *client.Client) *Service {
	return &Service{client}
}

func (s *Service) GetPrivatRate() (float64, error) {
	//use service client to execute http request and parse current rate
	s.client.PrivatBankClient.ParseRate()
	return float64(0), nil
}
