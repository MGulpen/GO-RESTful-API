package repository

import (

	//github.com/go-sql-driver/mysql is needed as anominous mysql driver plugin
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"GO-RESTful-API/model/vehicle/entity"
)

type IVehicleRepository interface {
	GetVehicleByLicensePlate(ctx context.Context, id string) (*entity.Vehicle, error)
	GetVehicles(ctx context.Context) ([]*entity.Vehicle, error)

	CreateVehicle()
	UpdateVehicle()
	DeleteVehicle()
}

type mysqlVehicleRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
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
	query := "SELECT * FROM cars WHERE license_plate = ?"

	list, err := db.fetch(ctx, query, licencePlate)
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

	query := "SELECT * FROM cars"

	list, err := db.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return list, err

}

func (db *mysqlVehicleRepository) CreateVehicle() {
	fmt.Println("Created new vehicle")
}

func (db *mysqlVehicleRepository) UpdateVehicle() {
	fmt.Println("Updated vehicle")
}

func (db *mysqlVehicleRepository) DeleteVehicle() {
	fmt.Println("Deleted vehicle")
}
