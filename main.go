package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	// register path on a DefaultServeMux
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Invalid data in request!", http.StatusBadRequest)
			return
		}
		_, _ = fmt.Fprintf(writer, "Hello %q", data)
	})
	http.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("GoodBye Hello World")
	})
	_ = http.ListenAndServe(":9090", nil)
}
