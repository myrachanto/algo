package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// see description (1)
type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

func CreateServer() {
	// see description (3)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root"))
	})

	// see description (4)
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		// see description (5)
		switch r.Method {
		case http.MethodGet:
			getProduct(w, r)
		case http.MethodPost:
			createProduct(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	// see description (6)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Cannot start server", err.Error())
		return
	}
}

// see description (7)
func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("On GET products")

	response := []Product{
		{
			ID:       "some-id-001",
			Name:     "PC Gamer",
			Price:    3500,
			Quantity: 10,
		},
		{
			ID:       "some-id-002",
			Name:     "PC Gamer Pro",
			Price:    4500,
			Quantity: 4,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// see description (8)
func createProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("On POST products")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var newProduct Product
	err = json.Unmarshal(body, &newProduct)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if newProduct.Name == "" {
		http.Error(w, "Bad Request - Name cannot be empty", http.StatusBadRequest)
		return
	}

	if newProduct.Price == 0 {
		http.Error(w, "Bad Request - Name cannot be $0", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
