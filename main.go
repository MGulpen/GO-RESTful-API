/*
*The main.go file is used to start the RESTful API.
 */
package main

import (
	serverSettings "GO-RESTful-API/config/server/settings"
	carsysHandler "GO-RESTful-API/controller/carsys/handler"
	rootHandler "GO-RESTful-API/controller/root/handler"
	vehicleHandler "GO-RESTful-API/controller/vehicle/handler"

	_ "github.com/go-sql-driver/mysql"

	//vehicleRepository "GO-RESTful-API/model/vehicle/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello from API")

	dbConn, err := sql.Open(`mysql`, "root:1143@tcp(localhost:3306)/vehicle_db")
	if err != nil {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		//os.Exit(1) docker?
	}

	defer dbConn.Close()

	//authorRepo := vehicleRepository.NewMysqlVehicleRepository(dbConn)
	http.HandleFunc("/", rootHandler.Agent)
	http.HandleFunc("/carsys/current-version", carsysHandler.CurrentVersion)
	http.HandleFunc("/objects/vehicle", vehicleHandler.Agent)
	http.ListenAndServe(serverSettings.Port, nil)

}
