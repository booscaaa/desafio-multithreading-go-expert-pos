package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Address struct {
	Data map[string]string
	API  string
}

var (
	BASE_URL_VIA_CEP = "http://viacep.com.br"
	// BASE_URL_API_CEP = "https://cdn.apicep.com"
	BASE_URL_API_CEP = "https://brasilapi.com.br/api/cep/v1"
)

func main() {
	timeout := time.Second * 1
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cep := "99150000"

	var wg sync.WaitGroup
	result := make(chan Address, 2)

	wg.Add(2)

	go getAddress(ctx, fmt.Sprintf("%s/%s", BASE_URL_API_CEP, cep), result, &wg)
	go getAddress(ctx, fmt.Sprintf("%s/ws/%s/json/", BASE_URL_VIA_CEP, cep), result, &wg)

	wg.Wait()
	close(result)

	select {
	case result := <-result:
		cancel()
		if len(result.Data) > 0 {
			fmt.Printf("API mais rápida: %s\n", result.API)
			fmt.Printf("Os dados são: %s\n", result.Data)
		} else {
			fmt.Println("Timeout: Nenhuma API respondeu dentro do prazo de 1 segundo.")
		}
	case <-time.After(timeout):
		fmt.Println("Timeout: Nenhuma API respondeu dentro do prazo de 1 segundo.")
	}
}

func getAddress(ctx context.Context, apiURL string, result chan<- Address, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, bytes.NewReader(nil))
	if err != nil {
		result <- Address{API: apiURL}
		return
	}
	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		result <- Address{API: apiURL}
		return
	}
	defer resp.Body.Close()

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		result <- Address{API: apiURL}
		return
	}

	result <- Address{API: apiURL, Data: data}
}
