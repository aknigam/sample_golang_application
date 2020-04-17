package service

import (
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
)

type ProductService struct {
	ProductRepo repository.ProductRepository
}

func (service *ProductService) GetProduct(productId int) (product *models.Product, appErr *errors.AppError) {
	product, err := service.ProductRepo.GetProduct(productId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "product not found", Code: -1}
		return
	}
	return
}

func (service *ProductService) DeleteProduct(productId int) (appErr *errors.AppError) {
	_, appErr = service.GetProduct(productId)
	if appErr != nil {
		return appErr
	}
	service.ProductRepo.DeleteProduct(productId)

	return
}

func (service *ProductService) UpdateProduct(product *models.Product) (appErr *errors.AppError) {
	_, err := service.GetProduct(product.Id)
	if err != nil {
		return err
	}
	productUpdateErr := service.ProductRepo.UpdateProduct(product)
	if productUpdateErr != nil {
		appErr = &errors.AppError{Error: productUpdateErr, Message: "product could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *ProductService) CreateProduct(product *models.Product) (productId int, appErr *errors.AppError) {
	productId, err := service.ProductRepo.CreateProduct(product)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "product could not be created", Code: -1}
	}
	return productId, appErr
}
