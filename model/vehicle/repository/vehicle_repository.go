package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	//github.com/go-sql-driver/mysql is needed as anominous mysql driver plugin
	_ "github.com/go-sql-driver/mysql"

	"GO-RESTful-API/model/vehicle/entity"
)

type IVehicleRepository interface {
	GetVehicleByLicensePlate(ctx context.Context, id string) (*entity.Vehicle, error)
	GetVehicles(ctx context.Context) ([]*entity.Vehicle, error)

	CreateVehicle(ctx context.Context, vehicle *entity.Vehicle) error
	UpdateVehicle(ctx context.Context, vehicle *entity.Vehicle) error
	DeleteVehicle(ctx context.Context, id string) error
}

type mysqlVehicleRepository struct {
	Conn *sql.DB
}

// NewMysqlVehicleRepository will create an object that represent the IVehicleRepository interface
func NewMysqlVehicleRepository(Conn *sql.DB) IVehicleRepository {

	return &mysqlVehicleRepository{Conn}
}

//fetch is a generic db call function to fetch the corrisponding objects.
func (db *mysqlVehicleRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.Vehicle, error) {

	rows, err := db.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		println(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*entity.Vehicle, 0)
	for rows.Next() {
		t := new(entity.Vehicle)
		err = rows.Scan(
			&t.LicensePlate,
			&t.Brand,
			&t.Model,
			&t.BuildDate,
			&t.OdometerValue,
			&t.OdometerType,
			&t.Transmission,
			&t.EngineType,
			&t.Retired,
		)

		if err != nil {
			println(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

//GetVehicle returns a vehicle record from db by license plate
func (db *mysqlVehicleRepository) GetVehicleByLicensePlate(ctx context.Context, licencePlate string) (*entity.Vehicle, error) {
	query := "SELECT * FROM cars WHERE license_plate = ? AND retired != ?"

	list, err := db.fetch(ctx, query, licencePlate, "yes")
	if err != nil {
		return nil, err
	}

	vehicle := &entity.Vehicle{}
	if len(list) > 0 {
		vehicle = list[0]
	} else {
		return nil, errors.New("Internal Server Error") //models.ErrNotFound
	}

	return vehicle, nil
}

//GetVehicles returns a slice of entity vehicles from db
func (db *mysqlVehicleRepository) GetVehicles(ctx context.Context) ([]*entity.Vehicle, error) {

	query := "SELECT * FROM cars WHERE retired != ?"

	list, err := db.fetch(ctx, query, "yes")
	if err != nil {
		return nil, err
	}

	return list, err

}

//CreateVehicle stores a new vehicle record into the db.
func (db *mysqlVehicleRepository) CreateVehicle(ctx context.Context, vehicle *entity.Vehicle) error {
	fmt.Println("Created new vehicle")

	query := `INSERT cars SET license_plate=? , brand=? , model=?, build_date=? , odometer_value=?, odometer_type=?, transmission=?, engine_type=?, retired=?`
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {

		return err
	}

	res, err := stmt.ExecContext(ctx, vehicle.LicensePlate, vehicle.Brand, vehicle.Model, vehicle.BuildDate, vehicle.OdometerValue, vehicle.OdometerType, vehicle.Transmission, vehicle.EngineType, vehicle.Retired)
	if err != nil {

		return err
	}
	fmt.Println(res)
	return nil
}

//UpdateVehicle updates an existing vehicle record
func (db *mysqlVehicleRepository) UpdateVehicle(ctx context.Context, vehicle *entity.Vehicle) error {
	fmt.Println("Updated vehicle")
	query := `UPDATE cars set brand=? , model=?, build_date=? , odometer_value=?, odometer_type=?, transmission=?, engine_type=? WHERE license_plate = ? AND retired != ?`

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, vehicle.Brand, vehicle.Model, vehicle.BuildDate, vehicle.OdometerValue, vehicle.OdometerType, vehicle.Transmission, vehicle.EngineType, vehicle.LicensePlate, "yes")
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return err
	}

	return nil
}

//DeleteVehicle soft deletes an existing vehicle record
func (db *mysqlVehicleRepository) DeleteVehicle(ctx context.Context, licencePlate string) error {
	fmt.Println("Deleted vehicle")
	query := `UPDATE cars set retired = ? WHERE license_plate = ?`

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, "yes", licencePlate)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return err
	}

	return nil
}
