package app

import (
	"fmt"
	"io"
	"log"

	"github.com/Windmill787/currency-parser/internal/client"
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
	fmt.Println("Client")

	client := client.NetClient()

	resp, err := client.Get("http://vseosvita.ua")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))
	fmt.Println("Response status: ", resp.StatusCode)
}
