package handler

import (
	"net/http"
	"text/template"
)

//Agent shows the index page.
func Agent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
