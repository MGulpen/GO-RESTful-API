package validator_test

import (
	vehicleVali "GO-RESTful-API/controller/vehicle/validator"
	"testing"
)

func Test_vehicleValidator_ValidateLicensePlate(t *testing.T) {
	type args struct {
		licensePlate string
	}
	tests := []struct {
		name      string
		validator vehicleVali.IVehicleValidator
		args      args
		want      bool
	}{
		{name: "test1", validator: vehicleVali.NewVehicleValidator(), args: args{licensePlate: "9-ABC-87"}, want: true},
		{name: "test2", validator: vehicleVali.NewVehicleValidator(), args: args{licensePlate: "9A-BC-87"}, want: true},
		{name: "test3", validator: vehicleVali.NewVehicleValidator(), args: args{licensePlate: "-9ABC-87"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.ValidateLicensePlate(tt.args.licensePlate); got != tt.want {
				t.Errorf("vehicleValidator.ValidateLicensePlate() = %v, want %v", got, tt.want)
			}
		})
	}
}
