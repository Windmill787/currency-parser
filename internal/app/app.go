package app

import (
	"github.com/Windmill787/currency-parser/internal/client"
	"github.com/Windmill787/currency-parser/internal/service"
)

func NewApp() {
	//create http client
	//create service that depends on client
	//use service to execute requests using http client

	//parse currensy rate for UAH -> USD
	//notify used to telegram
	//save this current rate to ram
	//set timeout for 1 hour
	//parse again
	//if rate is changed notify client, add the change amount to message
	client := client.NewClient()
	service := service.NewService(client)

	service.GetPrivatRate()
}
