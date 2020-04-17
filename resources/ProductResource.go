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

type ProductResource struct {
	productService service.ProductService
}

var Product ProductResource

func init() {
	Product := ProductResource{}
	Product.productService = service.ProductService{ProductRepo: repository.ProductRepository{}}
}

func (resource *ProductResource) CreateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	productId, _ := strconv.Atoi(c.Param("productId"))
	productId, err := resource.productService.CreateProduct(&product)
	if err != nil {
		common.Error.Println("Product could not be created", err)
		//return err
	}
	common.Info.Println("Created product with id ", productId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": product.Id})
	return

}

func (resource *ProductResource) GetProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	product, _ := resource.productService.GetProduct(productId)
	common.Info.Println("product found %d", productId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": product})
	return
}

func (resource *ProductResource) UpdateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	productId, _ := strconv.Atoi(c.Param("productId"))
	product.Id = productId
	err := resource.productService.UpdateProduct(&product)

	if err != nil {
		common.Error.Println("Product could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product updated successfully!"})
	return
}

func (resource *ProductResource) DeleteProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	err := resource.productService.DeleteProduct(productId)
	if err != nil {
		common.Error.Println("Product could not be deleted ", productId)
		//return err
	}
	common.Info.Println("Product deleted ", productId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}
