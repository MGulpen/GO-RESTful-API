package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":12345", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page\n")
}
