package service

import (
	"github.com/Windmill787/currency-parser/client"
	"github.com/Windmill787/currency-parser/entities"
)

type Service struct {
	client *client.Client
}

func NewService(client *client.Client) *Service {
	return &Service{client}
}

func (s *Service) GetPrivatRate(currency *entities.Currency) (float64, error) {
	return s.client.PrivatBankClient.ParseRate(currency)
}

func (s *Service) GetMonoRate(currency *entities.Currency) (float64, error) {
	return s.client.MonoBankClient.ParseRate(currency)
}
