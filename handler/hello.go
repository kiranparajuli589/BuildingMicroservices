package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	// read from request body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid data in request!", http.StatusBadRequest)
		return
	}
	// Now write to the response
	_, _ = fmt.Fprintf(w, "Hello %q", data)
}
