package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
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

func main() {
	for _, zipCode := range os.Args[1:] {
		resp, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when making request: %v\n", err)
		}
		defer resp.Body.Close()

		result, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(result, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing the response: %v\n", err)
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Zip Code: %s, City: %s, State: %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("File created successfully!")
		fmt.Println("City:", data.Localidade)
	}
}
