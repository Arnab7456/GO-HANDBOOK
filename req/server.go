package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

var products = []Product{
	{
		ID:          1,
		Name:        "Laptop Pro",
		Price:       999.99,
		Description: "High-performance laptop for professionals",
		Category:    "Electronics",
	},
	{
		ID:          2,
		Name:        "Wireless Headphones",
		Price:       199.99,
		Description: "Premium wireless headphones with noise cancellation",
		Category:    "Audio",
	},
	{
		ID:          3,
		Name:        "Smart Watch",
		Price:       299.99,
		Description: "Feature-rich smartwatch with health tracking",
		Category:    "Wearables",
	},
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", enableCORS(getProducts))

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
