package handler

// vehicleEntity "GO-RESTful-API/model/vehicle/entity"
// vehicleFactory "GO-RESTful-API/model/vehicle/factory"
// vehicleRepository "GO-RESTful-API/model/vehicle/repository"
import (
	//i_vehicleRepository "GO-RESTful-API/model/vehicle/repository/i_repository"
	//vehicleEntity "GO-RESTful-API/model/vehicle/entity"
	//"GO-RESTful-API/model/vehicle/repository"
	//"GO-RESTful-API/model/vehicle/repository"
	"net/http"
)

//Agent func delegates the corresponding method
func Agent(w http.ResponseWriter, r *http.Request) {
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
		if true {
			// use in if:a validator to check for valid car sign ( a func that checks with regex and then return true or false)

			//use function for retreiving vehicle by car lisence plate.
			//var vehicle entity.Vehicle = vehicleRepository.GetVehicle("43-PN-JK") // test
			//fmt.Println(vehicle)
		} else {
			//use function to return all vehicles.
			//var vehicles vehicleEntity = vehicleRepository.GetVehicles()
			//fmt.Println(vehicles[0].LicensePlate)
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
