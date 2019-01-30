package controller

import (
	serverSettings "GO-RESTful-API/config/server/settings"
	"net/http"

	carsysHandler "GO-RESTful-API/controller/carsys/handler"
	rootHandler "GO-RESTful-API/controller/root/handler"
	vehicleHandler "GO-RESTful-API/controller/vehicle/handler"
)

func Route() {
	//Add handlers.
	http.HandleFunc("/", rootHandler.Agent)
	http.HandleFunc("/carsys/current-version", carsysHandler.CurrentVersion)
	http.HandleFunc("/objects/vehicle", vehicleHandler.Agent)
	http.ListenAndServe(serverSettings.Port, nil)
}
