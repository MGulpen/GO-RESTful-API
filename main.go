/*
*The main.go file is used to start the RESTful API.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Message is a template for the json obj from http://nl.carsys.online/version.json
type Message struct {
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
	fmt.Fprintf(w, "Hello World\n")
}
func carsysCurrentVersionHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://nl.carsys.online/version.json")
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()
	var m Message

	//using io utils to read everything in the body.
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		// handle error
		panic(bodyErr)
	}
	//using json unmarshal to convert the incoming json object into an object
	jsonErr := json.Unmarshal(body, &m)
	if jsonErr != nil {
		// handle error
		panic(jsonErr)
	}

	//print the date/time stamp in the browser.
	fmt.Fprintf(w, m.BuildDate)
}
