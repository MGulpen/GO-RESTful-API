# GO-RESTful-API
A RESTful api made with GO

## How to use
I used the program Postman to interact with the API


* Show the date/time from carsys
```
GET: http://localhost:12345/carsys/current-version
```

* Return a single vehicle (json) by licenseplate :
```
GET: http://localhost:12345/objects/vehicle?licenseplate=ab-12-cd
```

* Return all vehicles (json):
```
GET: http://localhost:12345/objects/vehicle
```

* Store a single vehicle:
```
POST: http://localhost:12345/objects/vehicle

Set the body settings to "raw" and the "text" to "JSON(application/json)"
Url body example: 
    { "licenseplate":"88-AA-22",
	"brand":"someCarBrand",
	"model":"someCarModel",
	"build-date":"2001-01-01",
	"odometer-value":0,
	"odometer-type":"miles",
	"transmission":"manual",
	"engine-type":"2L"}
```

* Update a single vehicle:
```
PUT: http://localhost:12345/objects/vehicle?licenseplate=ab-12-cd

Set the body settings to "raw" and the "text" to "JSON(application/json)"
Url body example: 
    { "brand":"updatedCarBrand",
	"model":"updatedCarModel",
	"build-date":"2001-01-01",
	"odometer-value":0,
	"odometer-type":"miles",
	"transmission":"manual",
	"engine-type":"2L"}
```

* Delete a single vehicle by licenseplate (soft delete):
```
DELETE: http://localhost:12345/objects/vehicle?licenseplate=ab-12-cd
```