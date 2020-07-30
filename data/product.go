package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Products is a collection of Product
type Products []*Product

/*
	ToJSON serializes the contents of the collection to JSON
	NewEncoder provides better performance than json.Unmarshal as it does not
	have to buffer the output into an in memory slice of bytes
	This reduces allocations and the overheads of the service

	https://golang.org/pkg/encoding/json#NewEncoder
 */
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// returns list of products
func GetProducts() Products {
	return productsList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productsList = append(productsList, p)
}

func getNextId() int {
	lp := productsList[len(productsList) -1]
	return lp.ID + 1
}

func UpdateProduct(id int, product *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = id
	productsList[pos] = product
	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found.")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productsList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

// productList is a hard coded list of products for this example data source
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
