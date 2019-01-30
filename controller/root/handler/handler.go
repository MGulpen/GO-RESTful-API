package handler

import (
	"net/http"
	"text/template"
)

//Agent redirects to the corresponding method of the url header.
func Agent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
