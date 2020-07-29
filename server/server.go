package main

import (
	"context"
	"github.com/kiranparajuli589/building-microservices/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "building-microservices", log.LstdFlags)

	// create the handlers
	newHello := handler.NewHello(l)
	newBye := handler.NewBye(l)
	newProduct := handler.NewProduct(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/hello", newHello)
	sm.Handle("/bye", newBye)
	sm.Handle("/", newProduct)

	// create a new server
	s := &http.Server{
		Addr:              ":9090",           // configure the bind address
		Handler:           sm,                // set the default handler
		ErrorLog:          l,                 // set the logger for the server
		ReadHeaderTimeout: 1 * time.Second,   // max time to read request from the client
		WriteTimeout:      1 * time.Second,   // max time to write response to the client
		IdleTimeout:       120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Printf("Received %v order! Performing graceful SHUTDOWN...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(tc)
}
