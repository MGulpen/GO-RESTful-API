package handler

// vehicleEntity "GO-RESTful-API/model/vehicle/entity"
// vehicleFactory "GO-RESTful-API/model/vehicle/factory"
// vehicleRepository "GO-RESTful-API/model/vehicle/repository"
import (
	//i_vehicleRepository "GO-RESTful-API/model/vehicle/repository/i_repository"
	//vehicleEntity "GO-RESTful-API/model/vehicle/entity"
	//"GO-RESTful-API/model/vehicle/repository"
	//"GO-RESTful-API/model/vehicle/repository"

	contract "GO-RESTful-API/controller/vehicle/contract"
	vehicleVali "GO-RESTful-API/controller/vehicle/validator"
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

//Agent func delegates the corresponding method
func Agent(w http.ResponseWriter, r *http.Request) {
	dbConn, err := sql.Open(`mysql`, "root:1143@tcp(localhost:3306)/vehicle_db")
	if err != nil {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}

	defer dbConn.Close()
	repo := vehicleRepo.NewMysqlVehicleRepository(dbConn)
	c := context.Background()
	timeoutContext := time.Duration(5) * time.Second
	ctx, cancel := context.WithTimeout(c, timeoutContext)
	defer cancel()

	validator := vehicleVali.NewVehicleValidator()

	switch r.Method {
	case "POST":
		// add validator if json file is correct.
		vehicle := ConvertJsonToContract(r)
		if validator.ValidateVehicle(vehicle) {
			// create new vehicle
			fmt.Println("success")
			//vehicleRepository.CreateVehicle()
		} else {
			//show message couldn't create new vehicle.
		}

	case "GET":
		urlLicensePlateParam := r.URL.Query()["licenseplate"]
		if urlLicensePlateParam != nil {
			list, err := repo.GetVehicleByLicensePlate(ctx, urlLicensePlateParam[0])
			if err != nil {
				fmt.Println(err)
			}
			jsonObj, err := json.Marshal(list)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(w, string(jsonObj))

		} else {
			//use function to return all vehicles.

			list, err := repo.GetVehicles(ctx)

			if err != nil {
				fmt.Println(err)
			}

			jsonObj, err := json.Marshal(list)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(w, string(jsonObj))

		}
	case "PUT":
		// add validator if json file is correct. and response from db is not null with licenseplate.
		if true {
			// update existing  vehicle
			//vehicleRepository.UpdateVehicle()
		} else {
			//show message couldn't update vehicle. maybe validator message?
		}

	case "DELETE":
		// add validator if json file is correct.
		if true {
			// delete new vehicle
			//vehicleRepository.DeleteVehicle()

		} else {
			//show message couldn't delete vehicle.
		}
	default:
		http.Error(w, "Error, invalid request method - only supporting: GET, POST, PUT and DELETE", http.StatusNotImplemented)
	}
}

//CurrentVersion presents the date/time stamp from http://nl.carsys.online/version.json
func ConvertJsonToContract(r *http.Request) *contract.VehicleContract {

	//the defer tag forces the function to be executed at the end of the function.
	//defer resp.Body.Close()

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
	//print the date/time stamp in the browser.
	//fmt.Fprintf(w, carsysResponse.BuildDate)
}
