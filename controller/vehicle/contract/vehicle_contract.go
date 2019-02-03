package contract

//VehicleContract is used to communicate with clients
type VehicleContract struct {
	LicensePlate  string `json:"licenseplate"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	BuildDate     string `json:"build-date"`
	OdometerValue uint   `json:"odometer-value"`
	OdometerType  string `json:"odometer-type"`
	Transmission  string `json:"transmission"`
	EngineType    string `json:"engine-type"`
	//Retired       string `json:"retired"` //retired is not needed for clients its something used internaly
}
