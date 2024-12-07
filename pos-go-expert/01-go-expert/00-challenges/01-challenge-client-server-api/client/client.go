package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type DollarRate struct {
	Bid string `json:"bid"`
}

func main() {
	log.Println("main: REQUEST STARTED")
	defer log.Println("main: REQUEST COMPLETED")

	// 300ms context creation
	log.Println("main: 300ms context creation")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	// creation of the request using the created context
	log.Println("main: creation of the request using the created context")
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/dollar_rate", nil)
	HandleError(err)

	// performs the request using the created request
	log.Println("main: performs the request using the created request")
	resp, err := http.DefaultClient.Do(req)
	HandleError(err)
	defer resp.Body.Close()

	// read the body response
	log.Println("main: read the body response")
	body, err := io.ReadAll(resp.Body)
	HandleError(err)

	if resp.StatusCode == 500 {
		HandleError(errors.New(string(body)))
	}

	// performs json deserialization
	log.Println("main: performs json deserialization")
	var dollarRate DollarRate
	err = json.Unmarshal(body, &dollarRate)
	HandleError(err)

	// saves the dollar rate {dollarRate.Bid} in the file 'dollar_rate.txt'
	log.Println("main: saves the dollar rate '" + dollarRate.Bid + "' in the file 'dollar_rate.txt'")
	err = SaveDollarRateInFile(dollarRate.Bid)
	HandleError(err)

	// process completed successfully
	log.Println("main: process completed successfully")
}

func SaveDollarRateInFile(dollarRate string) error {
	// creation of the file 'dollar_rate.txt'
	log.Println("save_dollar_rate_in_file: creation of the file 'dollar_rate.txt'")
	file, err := os.Create("dollar_rate.txt")
	HandleError(err)

	// saves the dollar rate {dollarRate} in the file 'dollar_rate.txt'
	log.Println("save_dollar_rate_in_file: saves the dollar rate '" + dollarRate + "' in the file 'dollar_rate.txt'")
	_, err = file.Write([]byte("Dollar: " + dollarRate))
	HandleError(err)

	return nil
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
