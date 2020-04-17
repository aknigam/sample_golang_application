package service

import (
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
)

type PersonService struct {
	PersonRepo repository.PersonRepository
}

func (service *PersonService) GetPerson(personId int) (person *models.Person, appErr *errors.AppError) {
	person, err := service.PersonRepo.GetPerson(personId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "person not found", Code: -1}
		return
	}
	return
}

func (service *PersonService) DeletePerson(personId int) (appErr *errors.AppError) {
	_, appErr = service.GetPerson(personId)
	if appErr != nil {
		return appErr
	}
	service.PersonRepo.DeletePerson(personId)

	return
}

func (service *PersonService) UpdatePerson(person *models.Person) (appErr *errors.AppError) {
	_, err := service.GetPerson(person.Id)
	if err != nil {
		return err
	}
	personUpdateErr := service.PersonRepo.UpdatePerson(person)
	if personUpdateErr != nil {
		appErr = &errors.AppError{Error: personUpdateErr, Message: "person could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *PersonService) CreatePerson(person *models.Person) (personId int, appErr *errors.AppError) {
	personId, err := service.PersonRepo.CreatePerson(person)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "person could not be created", Code: -1}
	}
	return personId, appErr
}
