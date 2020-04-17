package service

import (
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
)

type CustomerService struct {
	CustomerRepo repository.CustomerRepository
}

func (service *CustomerService) GetCustomer(customerId int) (customer *models.Customer, appErr *errors.AppError) {
	customer, err := service.CustomerRepo.GetCustomer(customerId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "customer not found", Code: -1}
		return
	}
	return
}

func (service *CustomerService) DeleteCustomer(customerId int) (appErr *errors.AppError) {
	_, appErr = service.GetCustomer(customerId)
	if appErr != nil {
		return appErr
	}
	service.CustomerRepo.DeleteCustomer(customerId)

	return
}

func (service *CustomerService) UpdateCustomer(customer *models.Customer) (appErr *errors.AppError) {
	_, err := service.GetCustomer(customer.Id)
	if err != nil {
		return err
	}
	customerUpdateErr := service.CustomerRepo.UpdateCustomer(customer)
	if customerUpdateErr != nil {
		appErr = &errors.AppError{Error: customerUpdateErr, Message: "customer could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *CustomerService) CreateCustomer(customer *models.Customer) (customerId int, appErr *errors.AppError) {
	customerId, err := service.CustomerRepo.CreateCustomer(customer)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "customer could not be created", Code: -1}
	}
	return customerId, appErr
}
