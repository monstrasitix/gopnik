package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Product struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var (
	PRODUCTS = []Product{
		{
			Id:    "1",
			Title: "Shirt",
		},
		{
			Id:    "2",
			Title: "Shirt",
		},
	}
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(PRODUCTS)
	if err != nil {
		log.Fatal("Failed to marshal product")
	}

	w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	prod := Product{}
	found := false

	for _, product := range PRODUCTS {
		if product.Id == id {
			found = true
			prod = product

			break
		}
	}

	if found {
		data, err := json.Marshal(prod)
		if err != nil {
			log.Fatal("Failed to marshal product")
		}

		w.Write(data)
	} else {
		w.Write([]byte("null"))
	}
}
