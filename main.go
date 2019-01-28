/*
*The main.go file is used to start the RESTful API.
 */
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//CarsysResponse is a struct for the json obj from http://nl.carsys.online/version.json
type CarsysResponse struct {
	BuiltFromBranch string `json:"Built from branch"`
	CommitID        string `json:"Commit id"`
	BuildDate       string `json:"Build date"`
	BuildNumber     string `json:"Build number"`
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/carsys/current-version", carsysCurrentVersionHandler)
	http.ListenAndServe(":12345", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("GO-RESTful-API/view/index.html")
	t.Execute(w, nil)
	//fmt.Fprint(w, "view/index.html")

}

//carsysCurrentVersionHandler presents the date/time stamp from http://nl.carsys.online/version.json
func carsysCurrentVersionHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://nl.carsys.online/version.json")
	if err != nil {
		// handle error
		panic(err)
	}

	//the defer tag forces the function to be executed at the end of the function.
	defer resp.Body.Close()

	//using io utils to read everything in the body.
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		// handle error
		panic(bodyErr)
	}
	//using json unmarshal to convert the incoming json object.
	var carsysResponse CarsysResponse
	jsonErr := json.Unmarshal(body, &carsysResponse)
	if jsonErr != nil {
		// handle error
		panic(jsonErr)
	}

	//print the date/time stamp in the browser.
	fmt.Fprintf(w, carsysResponse.BuildDate)
}

func getVehicleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func postVehicleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}
func deleteVehicleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}
