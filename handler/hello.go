package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	/*
		Dependency Injection
		- For faster unit tests
		- Clean Handler
		- No directly creating concrete objects inside handler (avoid if possible)
		- Benefits:
			- In test we can use logger with something else. No worries about concrete implementations
			- Sometimes we want to log to file or to std out (more configurable)
	*/
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
