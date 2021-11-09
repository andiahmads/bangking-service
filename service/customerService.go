package service

import (
	"github.com/andiahmads/bangking-service/domain"
	"github.com/andiahmads/bangking-service/dto"
	"github.com/andiahmads/bangking-service/helpers"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *helpers.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *helpers.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *helpers.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inActive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)

}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *helpers.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	statusText := "active"

	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      statusText,
	}
	if response.Status == "0" {
		statusText = "inActive"
	}

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
