package service

import (
	"Banking/domain"
	"Banking/dto"
	"Banking/err"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerRest, *err.AppError)
	GetCustomerById(string) (*dto.CustomerRest, *err.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerRest, *err.AppError) {
	customerDom, error := s.repo.FindAll(status)

	if error != nil {
		return nil, error
	}

	customerRes := domain.ConvertToCustomerResArray(customerDom)

	return customerRes, error
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerRest, *err.AppError) {
	customerDom, error := s.repo.ById(id)

	if error != nil {
		return nil, error
	}

	customerRes := customerDom.ConvertToCustomerRest()
	return customerRes, error
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
