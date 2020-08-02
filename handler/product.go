package handler

import (
	"github.com/kiranparajuli589/building-microservices/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.l.Println("(GET) Product")
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.l.Println("(POST) Product")
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println("(PUT) Product")
		// expect id from the uri
		reg := regexp.MustCompile(`/([0-9]+)`)
		grp := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(grp) != 1 {
			p.l.Println(rw, "Invalid URI: more than one ID.")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(grp[0]) != 2 {
			p.l.Println("Invalid URI: more than one capture group.")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := grp[0][1]
		id, _ := strconv.Atoi(idString)

		p.updateProduct(id, rw, r)
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	/*
		One way of serializing structs into JSON is with Marshaling
		```
		d, err := json.Marshal(pl)
		if err != nil {
			http.Error(rw, "Unable to marshal d into JSON.", http.StatusInternalServerError)
		}
		_, err = rw.Write(d)
		if err != nil {
			http.Error(rw, "Unable to write d into JSON.", http.StatusInternalServerError)
		}
		````

		Marshaling and Encoder in a sense does same thing i.e writes the JSON encoding of v to the stream
		But rather than returning slice of data and error, encoder is writing the output directly to the io stream
		Why Encoder?: Don't have to buffer any Memory, don't have to allocate memory for that data object. Marginally Faster
	*/

	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to write data into JSON.", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod %#v", prod)
	data.AddProduct(prod)
}

func (p *Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found.", http.StatusInternalServerError)
	}
}
