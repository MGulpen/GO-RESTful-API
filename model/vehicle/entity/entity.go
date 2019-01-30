package entity

//Vehicle struct is for db query
type Vehicle struct {
	LicensePlate  string
	Brand         string
	Model         string
	BuildDate     string
	OdometerValue uint
	OdometerType  string
	Transmission  string
	EngineType    string
	Retired       string
}
