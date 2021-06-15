package handler

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Server struct {
		templates *template.Template
		asset     fs.FS
		assetFs   *hash.fs
	}
)

func NewServer() (*mux.Router, error) {
	s := &Server{}
	r := mux.NewRouter()
	r.HandleFunc("/", s.getHome)
	r.HandleFunc("/alist", s.getArricleList).Methods("GET")
	r.HandleFunc("/aget/{cat}", s.getAritcleShow).Methods("POST")

	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../assets"))))
	return r, nil
}

// func lookUpTemplate()
// func parseTempates()
