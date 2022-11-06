package domain

import "Banking/err"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

type CustomerRepo interface {
	FindAll(string) ([]Customer, *err.AppError)
	ById(string) (*Customer, *err.AppError)
}
