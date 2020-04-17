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

type OrderResource struct {
	orderService service.OrderService
}

var Order OrderResource

func init() {
	Order := OrderResource{}
	Order.orderService = service.OrderService{OrderRepo: repository.OrderRepository{}}
}

func (resource *OrderResource) CreateOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)
	orderId, _ := strconv.Atoi(c.Param("orderId"))
	orderId, err := resource.orderService.CreateOrder(&order)
	if err != nil {
		common.Error.Println("Order could not be created", err)
		//return err
	}
	common.Info.Println("Created order with id ", orderId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": order.Id})
	return

}

func (resource *OrderResource) GetOrder(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("orderId"))
	order, _ := resource.orderService.GetOrder(orderId)
	common.Info.Println("order found %d", orderId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": order})
	return
}

func (resource *OrderResource) UpdateOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)
	orderId, _ := strconv.Atoi(c.Param("orderId"))
	order.Id = orderId
	err := resource.orderService.UpdateOrder(&order)

	if err != nil {
		common.Error.Println("Order could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Order updated successfully!"})
	return
}

func (resource *OrderResource) DeleteOrder(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("orderId"))
	err := resource.orderService.DeleteOrder(orderId)
	if err != nil {
		common.Error.Println("Order could not be deleted ", orderId)
		//return err
	}
	common.Info.Println("Order deleted ", orderId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}
