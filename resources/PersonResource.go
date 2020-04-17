package resources

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample_golang_application/common"
	"sample_golang_application/models"
	"sample_golang_application/repository"
	"sample_golang_application/service"
	"strconv"
)

type PersonResource struct {
	personService service.PersonService
}

var Person PersonResource

func init() {
	Person := PersonResource{}
	Person.personService = service.PersonService{PersonRepo: repository.PersonRepository{}}
}

func (resource *PersonResource) CreatePerson(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)
	personId, _ := strconv.Atoi(c.Param("personId"))
	personId, err := resource.personService.CreatePerson(&person)
	if err != nil {
		common.Error.Println("Person could not be created", err)
		//return err
	}
	common.Info.Println("Created person with id ", personId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": person.Id})
	return

}

func (resource *PersonResource) GetPerson(c *gin.Context) {
	personId, _ := strconv.Atoi(c.Param("personId"))
	person, _ := resource.personService.GetPerson(personId)
	common.Info.Println("person found %d", personId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": person})
	return
}

func (resource *PersonResource) UpdatePerson(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)
	personId, _ := strconv.Atoi(c.Param("personId"))
	person.Id = personId
	err := resource.personService.UpdatePerson(&person)

	if err != nil {
		common.Error.Println("Person could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Person updated successfully!"})
	return
}

func (resource *PersonResource) DeletePerson(c *gin.Context) {
	personId, _ := strconv.Atoi(c.Param("personId"))
	err := resource.personService.DeletePerson(personId)
	if err != nil {
		common.Error.Println("Person could not be deleted ", personId)
		//return err
	}
	common.Info.Println("Person deleted ", personId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}
