package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	mysqldb "GO-RESTful-API/config/mysqldb/settings"
	"GO-RESTful-API/model/vehicle/entity"
)

func GetVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Vehicle\n")
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Multiple Vehicles\n")

	db, err := sql.Open("mysql", mysqldb.ConnectionString)
	if err != nil {
	}
	defer db.Close()

	q := "SELECT * FROM cars"
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var vehicle entity.Vehicle //maybe creatind a slice array to stack the results from the database.

	for rows.Next() {
		err := rows.Scan(
			&vehicle.LicensePlate,
			&vehicle.Brand,
			&vehicle.Model,
			&vehicle.BuildDate,
			&vehicle.OdometerValue,
			&vehicle.OdometerType,
			&vehicle.Transmission,
			&vehicle.EngineType,
			&vehicle.Retired)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(vehicle.LicensePlate)
	}
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created new vehicle\n")
}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updated vehicle\n")
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted vehicle\n")
}
