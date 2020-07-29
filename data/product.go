package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product defines the structure for an API product

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productsList
}

var productsList = []*Product {
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "jfd589",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}
