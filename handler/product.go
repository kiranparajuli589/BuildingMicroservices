package handler

import (
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
	pl := data.GetProducts()
	//// serializing into JSON with Marshaling
	//d, err := json.Marshal(pl)
	//if err != nil {
	//	http.Error(rw, "Unable to marshal d into JSON.", http.StatusInternalServerError)
	//}
	//_, err = rw.Write(d)
	//if err != nil {
	//	http.Error(rw, "Unable to write d into JSON.", http.StatusInternalServerError)
	//}

	// Marshaling and Encoder in a sense does same thing i.e writes the JSON encoding of v to the stream
	// But rather than returning slice of data and error, encoder is writing the output directly to the io stream
	// Why Encoder?: Don't have to buffer any Memory, don't have to allocate memory for that data object. Marginally Faster

	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to write data into JSON.", http.StatusInternalServerError)
	}
}
