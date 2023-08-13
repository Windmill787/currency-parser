package main

import (
	"log"
	"time"

	"github.com/Windmill787/currency-parser/client"
	"github.com/Windmill787/currency-parser/entities"
	"github.com/Windmill787/currency-parser/service"
)

func main() {
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

	currency := entities.USD()
	rate, err := service.GetPrivatRate(currency)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[%s] Bank={privat} currency={%s} rate={%.2f}\n", time.Now().Format(time.ANSIC), currency.Code, rate)

	rate, err = service.GetMonoRate(currency)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[%s] Bank={mono  } currency={%s} rate={%.2f}\n", time.Now().Format(time.ANSIC), currency.Code, rate)
}
