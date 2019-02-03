package handler

import (
	carsysContract "GO-RESTful-API/controller/carsys/contract"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CurrentVersion presents the date/time stamp from http://nl.carsys.online/version.json
func CurrentVersion(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://nl.carsys.online/version.json")
	if err != nil {
		// handle error
		panic(err)
	}

	//the defer tag forces the function to be executed at the end of the function.
	defer resp.Body.Close()

	//using io utils to read everything in the body.
	body, bodyErr := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Body)
	if bodyErr != nil {
		// handle error
		panic(bodyErr)
	}
	//using json unmarshal to convert the incoming json object.
	var carsysResponse carsysContract.CurrentVersion
	jsonErr := json.Unmarshal(body, &carsysResponse)
	if jsonErr != nil {
		// handle error
		panic(jsonErr)
	}

	//print the date/time stamp in the browser.
	fmt.Fprintf(w, carsysResponse.BuildDate)
}
