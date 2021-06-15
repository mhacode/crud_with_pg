package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) getArricleList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Article list")

}

func (s *Server) getAritcleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aritcle = %v\n", vars["cat"])

}
