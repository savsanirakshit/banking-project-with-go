package service

import (
	"Banking/domain"
	"Banking/err"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *err.AppError)
	GetCustomerById(string) (*domain.Customer, *err.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *err.AppError) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *err.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
