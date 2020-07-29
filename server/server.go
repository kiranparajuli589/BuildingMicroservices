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
	newHello := handler.NewHello(l)
	newBye := handler.NewBye(l)
	newProduct := handler.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", newHello)
	sm.Handle("/bye", newBye)
	sm.Handle("/", newProduct)

	s := &http.Server{
		Addr:              ":9090",
		Handler:           sm,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Printf("Received %v order! Performing graceful SHUTDOWN...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(tc)
}
