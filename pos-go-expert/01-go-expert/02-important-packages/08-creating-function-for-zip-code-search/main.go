package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", ZipCodeSearchHandler)
	http.ListenAndServe(":8080", nil)
}

func ZipCodeSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	zipCodeParam := r.URL.Query().Get("zipcode")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	zipCode, err := SearchZipCode(zipCodeParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// result, err := json.Marshal(zipCode)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)
	json.NewEncoder(w).Encode(zipCode)
}

func SearchZipCode(zipCode string) (*ViaCep, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c ViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
