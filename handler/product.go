package handler

import (
	"encoding/json"
	"github.com/kiranparajuli589/building-microservices/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product  {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	productsList := data.GetProducts()
	// serializing into JSON
	d, err := json.Marshal(productsList)
	if err != nil {
		http.Error(rw, "Unable to marshal d into JSON.", http.StatusInternalServerError)
	}
	_, err = rw.Write(d)
	if err != nil {
		http.Error(rw, "Unable to write d into JSON.", http.StatusInternalServerError)
	}
}
