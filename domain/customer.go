package domain

import "github.com/andiahmads/bangking-service/helpers"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *helpers.AppError)
	ById(string) (*Customer, *helpers.AppError)
}
