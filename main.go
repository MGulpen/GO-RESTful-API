/*
*The main.go file is used to start the RESTful API.
 */
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//CarsysResponse is a struct for the json obj from http://nl.carsys.online/version.json
type CarsysResponse struct {
	BuiltFromBranch string `json:"Built from branch"`
	CommitID        string `json:"Commit id"`
	BuildDate       string `json:"Build date"`
	BuildNumber     string `json:"Build number"`
}

//Cars struct is for db query
type Cars struct {
	LicensePlate  string
	Brand         string
	Model         string
	BuildDate     string
	OdometerValue uint
	OdometerType  string
	Transmission  string
	EngineType    string
	Retired       string
}

func main() {
	db, err := sql.Open("mysql", "root:1143@tcp(localhost:3306)/vehicle_db")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	q := "SELECT * FROM cars"
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var cars Cars //maybe creatind a slice array to stack the results from the database.

	for rows.Next() {

		err := rows.Scan(
			&cars.LicensePlate,
			&cars.Brand,
			&cars.Model,
			&cars.BuildDate,
			&cars.OdometerValue,
			&cars.OdometerType,
			&cars.Transmission,
			&cars.EngineType,
			&cars.Retired)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(cars.LicensePlate)
	}
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

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
		http.Error(w, "Error, invalid request method - only supporting: GET, POST, PUT and DELETE", http.StatusNotImplemented)
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
