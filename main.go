/*
*The main.go file is used to start the RESTful API.
 */
package main

import (
	controller "GO-RESTful-API/controller"
	"fmt"
	//serverSettings "GO-RESTful-API/config/server/settings"
	//carsysHandler "GO-RESTful-API/controller/carsys/handler"
	//rootHandler "GO-RESTful-API/controller/root/handler"
	//vehicleHandler "GO-RESTful-API/controller/vehicle/handler"
)

func main() {
	fmt.Println("Hello from API")
	//fmt.Println(serverSettings.Port)
	controller.Route()

	//http.HandleFunc("/", rootHandler.Agent)
	//http.HandleFunc("/carsys/current-version", carsysHandler.CurrentVersion)
	//http.HandleFunc("/objects/vehicle", vehicleHandler.Agent)
	//http.ListenAndServe(serverSettings.Port, nil)
}
