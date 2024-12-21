package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Structures that represent API responses

// Structure for the ViaZipCode API response
type ViaZipCodeResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// Structure for the Brazil API response
type BrazilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {
	// HTTP Server Configuration
	http.HandleFunc("/", zipCodeHandler)
	http.ListenAndServe(":8080", nil)
}

// Function that handles the request to consult the ZIP CODE
func zipCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Channels for receiving API responses
	viaZipCodeChannel := make(chan ViaZipCodeResponse)
	brazilZipCodeChannel := make(chan BrazilApiResponse)

	// 1 second timeout
	timeout := time.After(time.Second * 1)

	// Checks if the URL path is correct
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Gets the zip code parameter of the query string
	zipCodeParam := r.URL.Query().Get("zipCode")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Parameter 'zipCode' is mandatory"))
		return
	}

	// Starts goroutines to fetch data from both APIs
	go fetchViaZipCodeData(zipCodeParam, viaZipCodeChannel)
	go fetchBrazilApiData(zipCodeParam, brazilZipCodeChannel)

	// We use select to expect the fastest response
	select {
	case viaZipCodeData := <-viaZipCodeChannel:
		// If ViaZipCode API response arrived first, use it
		responseMessage := fmt.Sprintf("API: %s | Address: %s", "ViaZipCodeAPI", viaZipCodeData.Logradouro+" - "+viaZipCodeData.Bairro+" - "+viaZipCodeData.Localidade+" ("+viaZipCodeData.Estado+")")
		log.Println(responseMessage)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseMessage))

	case brazilZipCodeData := <-brazilZipCodeChannel:
		// If Brazil API response arrived first, use it
		responseMessage := fmt.Sprintf("API: %s | Address: %s", "BrazilAPI", brazilZipCodeData.Street+" - "+brazilZipCodeData.Neighborhood+" - "+brazilZipCodeData.City+" ("+brazilZipCodeData.State+")")
		log.Println(responseMessage)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseMessage))

	case <-timeout:
		// If the 1 second timeout is reached
		errorMessage := "Error: Timeout (1s)"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte(errorMessage))
	}
}

// Function to fetch data from the ViaZipCode API
func fetchViaZipCodeData(zipCode string, resultChannel chan<- ViaZipCodeResponse) {
	// Performs the GET request to the ViaZipCode API
	resp, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil {
		log.Println("Error when making the GET request for the ViaZipCode API (fetchViaZipCodeData):", err.Error())
		return
	}
	defer resp.Body.Close()

	// Read the response body
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response (fetchViaZipCodeData):", err.Error())
		return
	}

	// Deserialize the JSON from the response
	var zipCodeData ViaZipCodeResponse
	err = json.Unmarshal(result, &zipCodeData)
	if err != nil {
		log.Println("Error deserialize JSON (fetchViaZipCodeData):", err.Error())
		return
	}

	// Sends the deserialized data to the channel
	resultChannel <- zipCodeData
}

// Function to fetch data from the BrazilAPI API
func fetchBrazilApiData(zipCode string, resultChannel chan<- BrazilApiResponse) {
	// Performs the GET request to the BrazilAPI API
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + zipCode)
	if err != nil {
		log.Println("Error when making the GET request for the BrazilAPI API (fetchBrazilApiData):", err.Error())
		return
	}
	defer resp.Body.Close()

	// Read the response body
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response (fetchBrazilApiData):", err.Error())
		return
	}

	// Deserialize the JSON from the response
	var zipCodeData BrazilApiResponse
	err = json.Unmarshal(result, &zipCodeData)
	if err != nil {
		log.Println("Error deserialize JSON (fetchBrazilApiData):", err.Error())
		return
	}

	// Sends the deserialized data to the channel
	resultChannel <- zipCodeData
}
