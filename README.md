# GO-RESTful-API
A RESTful api made with GO

## How to use
I used the program Postman to interact with the API


*shows the date/time from carsys
```
GET: http://localhost:12345/carsys/current-version
```
*returns a single vehicle (json) by licenseplate :
```
GET: http://localhost:12345//objects/vehicle?licenseplate=ab-12-cd
```
*returns all vehicles (json):
```
GET: http://localhost:12345//objects/vehicle
```
*stores a single vehicle:
```
POST: http://localhost:12345//objects/vehicle
i set the body settings to "raw" and the "text" to "JSON(application/json)"*
body example: 
    { "licenseplate":"88-AA-22",
	"brand":"someCarBrand",
	"model":"someCarModel",
	"build-date":"2001-01-01",
	"odometer-value":0,
	"odometer-type":"miles",
	"transmission":"manual",
	"engine-type":"2L"}
```
