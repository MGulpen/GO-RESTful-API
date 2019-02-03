package handler

import (
	contract "GO-RESTful-API/controller/vehicle/contract"
	vehicleVali "GO-RESTful-API/controller/vehicle/validator"
	"GO-RESTful-API/model/vehicle/entity"
	vehicleRepo "GO-RESTful-API/model/vehicle/repository"
	"io/ioutil"

	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Agent func executes method depending on the http GET, POST, PUT and DELETE methods
func Agent(w http.ResponseWriter, r *http.Request) {
	dbConn, err := sql.Open(`mysql`, "root:1143@tcp(localhost:3306)/vehicle_db")
	if err != nil {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()
	repo := vehicleRepo.NewMysqlVehicleRepository(dbConn)

	c := context.Background()
	timeoutContext := time.Duration(5) * time.Second
	ctx, cancel := context.WithTimeout(c, timeoutContext)
	defer cancel()

	validator := vehicleVali.NewVehicleValidator()

	//determine which http method: GET, POST, PUT or DELETE
	switch r.Method {
	case "POST":
		// convert json obj to vehicle contract
		vehicle := ConvertJSONToContract(r)

		//validate the vehicle contract obj if its correct
		if validator.ValidateVehicle(vehicle) {

			// create new entity.vehicle for vehicle repository
			vehicleEntity := &entity.Vehicle{
				LicensePlate:  vehicle.LicensePlate,
				Brand:         vehicle.Brand,
				Model:         vehicle.Model,
				BuildDate:     vehicle.BuildDate,
				OdometerValue: vehicle.OdometerValue,
				OdometerType:  vehicle.OdometerType,
				Transmission:  vehicle.Transmission,
				EngineType:    vehicle.EngineType,
				Retired:       "no",
			}

			// use repository to create new vehicle
			err := repo.CreateVehicle(ctx, vehicleEntity)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			//show message couldn't create new vehicle.
			http.Error(w, "CreateVehicle-JSON object incorrect", http.StatusUnsupportedMediaType)
		}

	case "GET":
		//get licenseplate value from url param
		urlLicensePlateParam := r.URL.Query()["licenseplate"]
		if urlLicensePlateParam != nil {
			//get a single vehicle from repository by licenseplate
			list, err := repo.GetVehicleByLicensePlate(ctx, urlLicensePlateParam[0])
			if err != nil {
				fmt.Println(err)
			}
			//convert vehicle to json obj
			jsonObj, err := json.Marshal(list)
			if err != nil {
				fmt.Println(err)
				return
			}

			//print json obj
			fmt.Fprintln(w, string(jsonObj))

		} else {
			//get all vehicles from repository.
			list, err := repo.GetVehicles(ctx)
			if err != nil {
				fmt.Println(err)
			}

			//convert vehicles to json obj
			jsonObj, err := json.Marshal(list)
			if err != nil {
				fmt.Println(err)
				return
			}

			//print json obj
			fmt.Fprintln(w, string(jsonObj))

		}
	case "PUT":
		//get licenseplate value from url param
		urlLicensePlateParam := r.URL.Query()["licenseplate"]
		// convert json obj to vehicle contract
		vehicle := ConvertJSONToContract(r)
		//add licenseplate value from url, to the vehicle contract obj
		vehicle.LicensePlate = urlLicensePlateParam[0]

		//validate licenseplate value from url and the vehicle contract obj
		if urlLicensePlateParam != nil && validator.ValidateVehicle(vehicle) {
			vehicleEntity := &entity.Vehicle{
				LicensePlate:  vehicle.LicensePlate,
				Brand:         vehicle.Brand,
				Model:         vehicle.Model,
				BuildDate:     vehicle.BuildDate,
				OdometerValue: vehicle.OdometerValue,
				OdometerType:  vehicle.OdometerType,
				Transmission:  vehicle.Transmission,
				EngineType:    vehicle.EngineType,
			}

			// use repository to update new vehicle
			err := repo.UpdateVehicle(ctx, vehicleEntity)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			//show message couldn't update vehicle. maybe validator message?
			http.Error(w, "UpdateVehicle-JSON object incorrect", http.StatusUnsupportedMediaType)
		}

	case "DELETE":
		//get licenseplate value from url param
		urlLicensePlateParam := r.URL.Query()["licenseplate"]

		//validate licenseplate value from url
		if urlLicensePlateParam != nil && validator.ValidateLicensePlate(urlLicensePlateParam[0]) {

			// use repository to update new vehicle
			repo.DeleteVehicle(ctx, urlLicensePlateParam[0])

		} else {
			//show message couldn't delete vehicle.
			http.Error(w, "DeleteVehicle-lisenceplate incorrect", http.StatusNotFound)

		}
	default:
		http.Error(w, "Error, invalid request method - only supporting: GET, POST, PUT and DELETE", http.StatusNotImplemented)
	}
}

//ConvertJSONToContract converts the json obj from the http.body to a vehicle contract obj
func ConvertJSONToContract(r *http.Request) *contract.VehicleContract {

	//using io utils to read everything in the body.
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		// handle error
		panic(bodyErr)
	}
	//using json unmarshal to convert the incoming json object.
	var vehicleContract contract.VehicleContract
	jsonErr := json.Unmarshal(body, &vehicleContract)
	if jsonErr != nil {
		// handle error
		panic(jsonErr)
	}
	return &vehicleContract
}
