package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const baseURL string = "http://localhost:8080"

type product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

// see description (1)
func GetProducts() {
	log.Println("Getting products")

	// see description (2)
	resp, err := http.Get(baseURL + "/products")
	if err != nil {
		log.Fatal("Error requesting products: %w", err)
	}

	// see description (3)
	var respBody []product
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body bytes response: %w", err)
	}

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Fatal("Error unmarshaling response: %w", err)
	}
	log.Println(respBody)
}

// see description (4)
func CreateProduct() {
	log.Println("Creating product")

	// see description (5)
	newProduct := product{
		Name:     "Gaming Console",
		Price:    120,
		Quantity: 3,
	}

	requestBody, err := json.Marshal(newProduct)
	if err != nil {
	}

	// see description (6)
	resp, err := http.Post(baseURL+"/products",
		"application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error requesting products: %w", err)
	}

	log.Println(resp.Status)
	log.Println(resp.StatusCode)
}
