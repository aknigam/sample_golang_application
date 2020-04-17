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

type SalesPersonResource struct {
	salesPersonService service.SalesPersonService
}

var SalesPerson SalesPersonResource

func init() {
	SalesPerson := SalesPersonResource{}
	SalesPerson.salesPersonService = service.SalesPersonService{SalesPersonRepo: repository.SalesPersonRepository{}}
}

func (resource *SalesPersonResource) CreateSalesPerson(c *gin.Context) {
	var salesPerson models.SalesPerson
	c.BindJSON(&salesPerson)
	salesPersonId, _ := strconv.Atoi(c.Param("salesPersonId"))
	salesPersonId, err := resource.salesPersonService.CreateSalesPerson(&salesPerson)
	if err != nil {
		common.Error.Println("SalesPerson could not be created", err)
		//return err
	}
	common.Info.Println("Created salesPerson with id ", salesPersonId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": salesPerson.Id})
	return

}

func (resource *SalesPersonResource) GetSalesPerson(c *gin.Context) {
	salesPersonId, _ := strconv.Atoi(c.Param("salesPersonId"))
	salesPerson, _ := resource.salesPersonService.GetSalesPerson(salesPersonId)
	common.Info.Println("salesPerson found %d", salesPersonId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": salesPerson})
	return
}

func (resource *SalesPersonResource) UpdateSalesPerson(c *gin.Context) {
	var salesPerson models.SalesPerson
	c.BindJSON(&salesPerson)
	salesPersonId, _ := strconv.Atoi(c.Param("salesPersonId"))
	salesPerson.Id = salesPersonId
	err := resource.salesPersonService.UpdateSalesPerson(&salesPerson)

	if err != nil {
		common.Error.Println("SalesPerson could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "SalesPerson updated successfully!"})
	return
}

func (resource *SalesPersonResource) DeleteSalesPerson(c *gin.Context) {
	salesPersonId, _ := strconv.Atoi(c.Param("salesPersonId"))
	err := resource.salesPersonService.DeleteSalesPerson(salesPersonId)
	if err != nil {
		common.Error.Println("SalesPerson could not be deleted ", salesPersonId)
		//return err
	}
	common.Info.Println("SalesPerson deleted ", salesPersonId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}
