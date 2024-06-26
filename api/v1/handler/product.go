package apiV1handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/monstrasitix/gopnik/model"
)

var (
	PRODUCTS = []model.Product{
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
		w.Write([]byte("null"))
	} else {
		w.Write(data)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	prod := model.Product{}
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
			w.Write([]byte("null"))
		} else {
			w.Write(data)
		}
	} else {
		w.Write([]byte("null"))
	}
}
