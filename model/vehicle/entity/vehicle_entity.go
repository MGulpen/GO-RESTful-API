package entity

//Vehicle struct is for db query
type Vehicle struct {
	LicensePlate  string `json:"license_plate"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	BuildDate     string `json:"build_date"`
	OdometerValue uint   `json:"odometer_value"`
	OdometerType  string `json:"odometer_type"`
	Transmission  string `json:"transmission"`
	EngineType    string `json:"engine_type"`
	Retired       string `json:"retired"`
}
