package handler

// vehicleEntity "GO-RESTful-API/model/vehicle/entity"
// vehicleFactory "GO-RESTful-API/model/vehicle/factory"
// vehicleRepository "GO-RESTful-API/model/vehicle/repository"
import (
	//i_vehicleRepository "GO-RESTful-API/model/vehicle/repository/i_repository"
	//vehicleEntity "GO-RESTful-API/model/vehicle/entity"
	//"GO-RESTful-API/model/vehicle/repository"
	//"GO-RESTful-API/model/vehicle/repository"

	vehicleRepo "GO-RESTful-API/model/vehicle/repository"
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

	switch r.Method {
	case "POST":
		// add validator if json file is correct.

		if true {
			// create new vehicle

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
