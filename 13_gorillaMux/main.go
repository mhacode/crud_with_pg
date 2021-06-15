package main

import (
	"crud_with_pg/13_gorillaMux/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	r, err := handler.NewServer()

	if err != nil {
		log.Fatal("Handler not found")
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
