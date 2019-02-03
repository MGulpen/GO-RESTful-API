package validator

import (
	obj "GO-RESTful-API/controller/vehicle/contract"
)

type IVehicleValidator interface {
	ValidateLicensePlate(id string) bool
	ValidateBrand(id string) bool
	ValidateModel(id string) bool
	ValidateBuildDate(id string) bool
	ValidateOdometerValue(id uint) bool
	ValidateOdometerType(id string) bool
	ValidateTransmission(id string) bool
	ValidateEngineType(id string) bool
	ValidateVehicle(vehicle *obj.VehicleContract) bool
}

type vehicleValidator struct {
}

// NewVehicleValidator will create an object that represent the IVehicleValidator interface
func NewVehicleValidator() IVehicleValidator {
	return &vehicleValidator{}
}

func (validator *vehicleValidator) ValidateLicensePlate(id string) bool {
	return true
}
func (validator *vehicleValidator) ValidateBrand(id string) bool {
	return true
}
func (validator *vehicleValidator) ValidateModel(id string) bool {
	return true
}
func (validator *vehicleValidator) ValidateBuildDate(id string) bool {
	return true
}
func (validator *vehicleValidator) ValidateOdometerValue(id uint) bool {
	return true
}
func (validator *vehicleValidator) ValidateOdometerType(id string) bool {
	return true
}
func (validator *vehicleValidator) ValidateTransmission(id string) bool {
	return true
}
func (validator *vehicleValidator) ValidateEngineType(id string) bool {
	return true
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
