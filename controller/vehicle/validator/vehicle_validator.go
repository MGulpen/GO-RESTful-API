package validator

import (
	obj "GO-RESTful-API/controller/vehicle/contract"
	"regexp"
)

type IVehicleValidator interface {
	ValidateLicensePlate(licensePlate string) bool
	ValidateBrand(brand string) bool
	ValidateModel(model string) bool
	ValidateBuildDate(buildDate string) bool
	ValidateOdometerValue(odometerValue uint) bool
	ValidateOdometerType(odometerType string) bool
	ValidateTransmission(transmission string) bool
	ValidateEngineType(engineType string) bool
	ValidateVehicle(vehicle *obj.VehicleContract) bool
}

//vehicleValidator is used as return obj that can use the validator methods
type vehicleValidator struct {
}

// NewVehicleValidator will create an object that represent the IVehicleValidator interface
func NewVehicleValidator() IVehicleValidator {
	return &vehicleValidator{}
}

//ValidateLicensePlate checks if the string input matches that of a vehicle licenseplate
func (validator *vehicleValidator) ValidateLicensePlate(licensePlate string) bool {
	var validID = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\-]*$`)
	isValid := validID.MatchString(licensePlate)
	return isValid
}

//ValidateBrand checks if the string input matches that of a vehicle brand
func (validator *vehicleValidator) ValidateBrand(brand string) bool {
	var validID = regexp.MustCompile(`^[A-Za-z][A-Za-z]*$`)
	isValid := validID.MatchString(brand)
	return isValid
}

//ValidateModel checks if the string input matches that of a vehicle model
func (validator *vehicleValidator) ValidateModel(model string) bool {
	var validID = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\-]*$`)
	isValid := validID.MatchString(model)
	return isValid
}

//ValidateBuildDate checks if the string input matches that of a vehicle build date
func (validator *vehicleValidator) ValidateBuildDate(buildDate string) bool {
	//return len(buildDate) > 0 //check for time format? 2001-01-01
	//[0-9]\{4\}\-[0-9]\{2\}\-[0-9]\{2\} doesnt work?
	var validID = regexp.MustCompile(`^[0-9][0-9][0-9][0-9]\-[0-9][0-9]\-[0-9][0-9]$`)
	isValid := validID.MatchString(buildDate)
	return isValid
}

//ValidateOdometerValue checks if the uint input is above zero
func (validator *vehicleValidator) ValidateOdometerValue(odometerValue uint) bool {
	return odometerValue >= 0
}

//ValidateOdometerType checks if the odometer type is of "miles" or "kilometers"
func (validator *vehicleValidator) ValidateOdometerType(odometerType string) bool {
	return (odometerType == "miles" || odometerType == "kilometers")
}

//ValidateTransmission checks if the transmission is "manual" or "automatic"
func (validator *vehicleValidator) ValidateTransmission(transmission string) bool {
	return (transmission == "manual" || transmission == "automatic")
}

//ValidateEngineType checks if the string input isn't empty
func (validator *vehicleValidator) ValidateEngineType(engineType string) bool {
	return len(engineType) > 0
}

//ValidateVehicle uses multiple validator methods to determine if vehicle obj is valid
func (validator *vehicleValidator) ValidateVehicle(vehicle *obj.VehicleContract) bool {

	return (validator.ValidateLicensePlate(vehicle.LicensePlate) &&
		validator.ValidateBrand(vehicle.Brand) &&
		validator.ValidateModel(vehicle.Model) &&
		validator.ValidateBuildDate(vehicle.BuildDate) &&
		validator.ValidateOdometerValue(vehicle.OdometerValue) &&
		validator.ValidateOdometerType(vehicle.OdometerType) &&
		validator.ValidateTransmission(vehicle.Transmission) &&
		validator.ValidateEngineType(vehicle.EngineType))
}
