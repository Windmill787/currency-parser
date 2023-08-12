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

func (s *Service) GetPrivatRate(currency string) (float64, error) {
	return s.client.PrivatBankClient.ParseRate(currency)
}
