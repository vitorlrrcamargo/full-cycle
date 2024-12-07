package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type DollarRate struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/dollar_rate", DollarRateHandler)
	http.ListenAndServe(":8080", nil)
}

func DollarRateHandler(w http.ResponseWriter, r *http.Request) {
	println()
	log.Println("dollar_rate_handler: REQUEST STARTED")
	defer log.Println("dollar_rate_handler: REQUEST COMPLETED")

	// search for the dollar rate in the api 'economia.awesomeapi.com.br'
	log.Println("dollar_rate_handler: search for the dollar rate in the api 'economia.awesomeapi.com.br'")
	dollarRate, err := SearchDollarRate()
	if err != nil {
		// error when executing the SearchDollarRate() function
		log.Println("ERRO: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// saves the dollar rate {dollarRate.Usdbrl.Bid} in the 'app.db' database
	log.Println("dollar_rate_handler: saves the dollar rate '" + dollarRate.Usdbrl.Bid + "' in the 'app.db' database")
	err = SaveDollarRateInDatabase(dollarRate.Usdbrl.Bid)
	if err != nil {
		// error when executing the SaveDollarRateInDatabase() function
		log.Println("ERRO: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// responds to the request with the dollar rate
	log.Println("dollar_rate_handler: responds to the request with the dollar rate")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"bid": dollarRate.Usdbrl.Bid})

	// process completed successfully
	log.Println("dollar_rate_handler: process completed successfully")
}

func SearchDollarRate() (*DollarRate, error) {
	// 200ms context creation
	log.Println("search_dollar_rate: 200ms context creation")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	// creation of the request using the created context
	log.Println("search_dollar_rate: creation of the request using the created context")
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/usd-brl", nil)
	if err != nil {
		return nil, err
	}

	// performs the request using the created request
	log.Println("search_dollar_rate: performs the request using the created request")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the body response
	log.Println("search_dollar_rate: read the body response")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// performs json deserialization
	log.Println("search_dollar_rate: performs json deserialization")
	var dollarRate DollarRate
	err = json.Unmarshal(body, &dollarRate)
	if err != nil {
		return nil, err
	}

	return &dollarRate, nil
}

func SaveDollarRateInDatabase(dollarRate string) error {
	// 10ms context creation
	log.Println("save_dollar_rate_in_database: 10ms context creation")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	// database initialization
	log.Println("save_dollar_rate_in_database: database initialization")
	db, err := sql.Open("sqlite", "app.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// creation of the 'dollar_rate' table if it does not exist
	log.Println("save_dollar_rate_in_database: creation of the 'dollar_rate' table if it does not exist")
	sqlCreateTable := `CREATE TABLE IF NOT EXISTS dollar_rate (dollar TEXT)`
	_, err = db.Exec(sqlCreateTable)
	if err != nil {
		return err
	}

	// prepares the query for insertion into the database
	log.Println("save_dollar_rate_in_database: prepares the query for insertion into the database")
	stmt, err := db.Prepare("INSERT INTO dollar_rate(dollar) VALUES(?)")
	if err != nil {
		return err
	}

	// runs the query to save the dollar rate {dollarRate} in the 'dollar_rate' table
	log.Println("save_dollar_rate_in_database: runs the query to save the dollar rate '" + dollarRate + "' in the 'dollar_rate' table")
	db.Exec("BEGIN TRANSACTION")
	_, err = stmt.ExecContext(ctx, dollarRate)
	if err != nil {
		return err
	}
	db.Exec("COMMIT")
	db.Exec("END TRANSACTION")

	return nil
}
