package domain

import (
	"Banking/dto"
	"Banking/err"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (receiver Customer) StatusAsText() string {
	status := "active"
	if receiver.Status == "0" {
		status = "inactive"
	}
	return status
}

type CustomerRepo interface {
	FindAll(string) ([]Customer, *err.AppError)
	ById(string) (*Customer, *err.AppError)
}

func (c Customer) ConvertToCustomerRest() (res *dto.CustomerRest) {
	return &dto.CustomerRest{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusAsText(),
	}
}

func ConvertToCustomerResArray(domainObj []Customer) (resObj []dto.CustomerRest) {
	customerRes := make([]dto.CustomerRest, 0)
	for _, customer := range domainObj {
		customerRes = append(customerRes, dto.CustomerRest{
			Id:          customer.Id,
			Name:        customer.Name,
			City:        customer.City,
			Zipcode:     customer.Zipcode,
			DateOfBirth: customer.DateOfBirth,
			Status:      customer.StatusAsText(),
		})
	}
	return customerRes
}
