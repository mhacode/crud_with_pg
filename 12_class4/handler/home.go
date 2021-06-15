package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	template := s.templates.Lookup("home.html")
	if template == nil {
		log.Println("unable to load exectut")
	}
	http.Redirect(w, r, http.StatusSeeOther)
	return


	if condition {
		
	}

}
