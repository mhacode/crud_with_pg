package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) getArticleList(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.New("articles.html").ParseFiles("./assets/templates/articles.html")

	err := tmp.Execute(w, nil)
	if err != nil {
		log.Println("Error Executing templates : ", err)
		return
	}

}

func (s *Server) getArticleShow(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprint(w, "Hello From Article Details", id)
}
