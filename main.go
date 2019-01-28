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
	http.HandleFunc("/carsys/current-version", carsysCurrentVersion)
	http.HandleFunc("/objects/vehicle", vehicleHandler)
	http.ListenAndServe(":12345", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

//carsysCurrentVersionHandler presents the date/time stamp from http://nl.carsys.online/version.json
func carsysCurrentVersion(w http.ResponseWriter, r *http.Request) {
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

func vehicleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// add validator if json file is correct.
		if true {
			// create new vehicle
			createVehicle(w, r)
		} else {
			//show message couldn't create new vehicle.
		}

	case "GET":
		if true {
			// use in if:a validator to check for valid car sign ( a func that checks with regex and then return true or false)

			//use function for retreiving vehicle by car lisence plate.
			getVehicle(w, r)
		} else {
			//use function to return all vehicles.
			getVehicles(w, r)
		}
	case "PUT":
		// add validator if json file is correct. and response from db is not null with licenseplate.
		if true {
			// update existing  vehicle
			updateVehicle(w, r)
		} else {
			//show message couldn't update vehicle. maybe validator message?
		}

	case "DELETE":
		// add validator if json file is correct.
		if true {
			// delete new vehicle
			deleteVehicle(w, r)

		} else {
			//show message couldn't delete vehicle.
		}
	default:
		fmt.Fprintf(w, "Error, invalid request method - only supporting: GET, POST, PUT and DELETE")
	}
}

func getVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Vehicle\n")
}

func getVehicles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Multiple Vehicles\n")
}

func createVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created new vehicle\n")
}

func updateVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updated vehicle\n")
}

func deleteVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted vehicle\n")
}
