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

type CustomerResource struct {
	customerService service.CustomerService
}

var Customer CustomerResource

func init() {
	Customer := CustomerResource{}
	Customer.customerService = service.CustomerService{CustomerRepo: repository.CustomerRepository{}}
}

func (resource *CustomerResource) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	customerId, _ := strconv.Atoi(c.Param("customerId"))
	customerId, err := resource.customerService.CreateCustomer(&customer)
	if err != nil {
		common.Error.Println("Customer could not be created", err)
		//return err
	}
	common.Info.Println("Created customer with id ", customerId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": customer.Id})
	return

}

func (resource *CustomerResource) GetCustomer(c *gin.Context) {
	customerId, _ := strconv.Atoi(c.Param("customerId"))
	customer, _ := resource.customerService.GetCustomer(customerId)
	common.Info.Println("customer found %d", customerId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": customer})
	return
}

func (resource *CustomerResource) UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	customerId, _ := strconv.Atoi(c.Param("customerId"))
	customer.Id = customerId
	err := resource.customerService.UpdateCustomer(&customer)

	if err != nil {
		common.Error.Println("Customer could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Customer updated successfully!"})
	return
}

func (resource *CustomerResource) DeleteCustomer(c *gin.Context) {
	customerId, _ := strconv.Atoi(c.Param("customerId"))
	err := resource.customerService.DeleteCustomer(customerId)
	if err != nil {
		common.Error.Println("Customer could not be deleted ", customerId)
		//return err
	}
	common.Info.Println("Customer deleted ", customerId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}
