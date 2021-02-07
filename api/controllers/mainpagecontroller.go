package controllers

import (
	"fmt"
	"net/http"
)

// Home wraps main page of API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", "Welcome to StatInfoAPI")

}
