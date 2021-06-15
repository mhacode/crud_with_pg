// Gorilla musk
// FuncMap
//path
//assetfs
// gorilla mux static files
package main

import (
	"crud_with_pg/12_class4/handler"
	"log"
	"net/http"
	"time"
)

func main() {

	r, err := handler.NewServer()
	if err != nil {
		log.Fatal("handler not found")
	}

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
