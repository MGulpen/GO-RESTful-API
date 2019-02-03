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

type vehicleValidator struct {
}

// NewVehicleValidator will create an object that represent the IVehicleValidator interface
func NewVehicleValidator() IVehicleValidator {
	return &vehicleValidator{}
}

func (validator *vehicleValidator) ValidateLicensePlate(licensePlate string) bool {
	var validID = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\-]*$`)
	isValid := validID.MatchString(licensePlate)
	return isValid
}
func (validator *vehicleValidator) ValidateBrand(brand string) bool {
	var validID = regexp.MustCompile(`^[A-Za-z][A-Za-z]*$`)
	isValid := validID.MatchString(brand)
	return isValid
}
func (validator *vehicleValidator) ValidateModel(model string) bool {
	var validID = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\-]*$`)
	isValid := validID.MatchString(model)
	return isValid
}
func (validator *vehicleValidator) ValidateBuildDate(buildDate string) bool {
	//return len(buildDate) > 0 //check for time format? 2001-01-01
	//[0-9]\{4\}\-[0-9]\{2\}\-[0-9]\{2\} doesnt work?
	var validID = regexp.MustCompile(`^[0-9][0-9][0-9][0-9]\-[0-9][0-9]\-[0-9][0-9]$`)
	isValid := validID.MatchString(buildDate)
	return isValid
}
func (validator *vehicleValidator) ValidateOdometerValue(odometerValue uint) bool {
	return odometerValue >= 0
}
func (validator *vehicleValidator) ValidateOdometerType(odometerType string) bool {
	return (odometerType == "miles" || odometerType == "kilometers")
}
func (validator *vehicleValidator) ValidateTransmission(transmission string) bool {
	return (transmission == "manual" || transmission == "automatic")
}
func (validator *vehicleValidator) ValidateEngineType(engineType string) bool {
	return len(engineType) > 0
}

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
