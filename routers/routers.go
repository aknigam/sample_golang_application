package routers

import (
	"github.com/gin-gonic/gin"
	"sample_golang_application/resources"
)

func SetApplicationRoutes(r *gin.RouterGroup) {

	r.GET("sprints/:sprintId", resources.Sprint.GetSprint)
	r.PUT("sprints/:sprintId", resources.Sprint.UpdateSprint)
	r.POST("sprints", resources.Sprint.CreateSprint)
	r.DELETE("sprints/:sprintId", resources.Sprint.DeleteSprint)

	r.GET("sprints/:sprintId/stories/:storyId", resources.Sprint.GetSprintStories)
	r.PUT("sprints/:sprintId/stories/:storyId", resources.Sprint.UpdateSprintStories)
	r.POST("sprints/:sprintId/stories", resources.Sprint.CreateSprintStories)
	r.DELETE("sprints/:sprintId/stories/:storyId", resources.Sprint.DeleteSprintStories)
	r.POST("sprints/:sprintId/stories/:storyId/link", resources.Sprint.LinkSprintStories)
	r.POST("sprints/:sprintId/stories/:storyId/unlink", resources.Sprint.UnlinkSprintStories)

	r.GET("storys/:storyId/tasks/:taskId", resources.Story.GetStoryTasks)
	r.PUT("storys/:storyId/tasks/:taskId", resources.Story.UpdateStoryTasks)
	r.POST("storys/:storyId/tasks", resources.Story.CreateStoryTasks)
	r.DELETE("storys/:storyId/tasks/:taskId", resources.Story.DeleteStoryTasks)
	r.POST("storys/:storyId/tasks/:taskId/link", resources.Story.LinkStoryTasks)
	r.POST("storys/:storyId/tasks/:taskId/unlink", resources.Story.UnlinkStoryTasks)

	r.GET("storys/:storyId/comments/:commentId", resources.Story.GetStoryComments)
	r.PUT("storys/:storyId/comments/:commentId", resources.Story.UpdateStoryComments)
	r.POST("storys/:storyId/comments", resources.Story.CreateStoryComments)
	r.DELETE("storys/:storyId/comments/:commentId", resources.Story.DeleteStoryComments)
	r.POST("storys/:storyId/comments/:commentId/link", resources.Story.LinkStoryComments)
	r.POST("storys/:storyId/comments/:commentId/unlink", resources.Story.UnlinkStoryComments)

	r.GET("storys/:storyId/poComments/:commentId", resources.Story.GetStoryPoComments)
	r.PUT("storys/:storyId/poComments/:commentId", resources.Story.UpdateStoryPoComments)
	r.POST("storys/:storyId/poComments", resources.Story.CreateStoryPoComments)
	r.DELETE("storys/:storyId/poComments/:commentId", resources.Story.DeleteStoryPoComments)
	r.POST("storys/:storyId/poComments/:commentId/link", resources.Story.LinkStoryPoComments)
	r.POST("storys/:storyId/poComments/:commentId/unlink", resources.Story.UnlinkStoryPoComments)

	r.GET("persons/:personId", resources.Person.GetPerson)
	r.PUT("persons/:personId", resources.Person.UpdatePerson)
	r.POST("persons", resources.Person.CreatePerson)
	r.DELETE("persons/:personId", resources.Person.DeletePerson)

	r.GET("products/:productId", resources.Product.GetProduct)
	r.PUT("products/:productId", resources.Product.UpdateProduct)
	r.POST("products", resources.Product.CreateProduct)
	r.DELETE("products/:productId", resources.Product.DeleteProduct)

	r.GET("orders/:orderId", resources.Order.GetOrder)
	r.PUT("orders/:orderId", resources.Order.UpdateOrder)
	r.POST("orders", resources.Order.CreateOrder)
	r.DELETE("orders/:orderId", resources.Order.DeleteOrder)

	r.GET("customers/:customerId", resources.Customer.GetCustomer)
	r.PUT("customers/:customerId", resources.Customer.UpdateCustomer)
	r.POST("customers", resources.Customer.CreateCustomer)
	r.DELETE("customers/:customerId", resources.Customer.DeleteCustomer)

	r.GET("salespersons/:salesPersonId", resources.SalesPerson.GetSalesPerson)
	r.PUT("salespersons/:salesPersonId", resources.SalesPerson.UpdateSalesPerson)
	r.POST("salespersons", resources.SalesPerson.CreateSalesPerson)
	r.DELETE("salespersons/:salesPersonId", resources.SalesPerson.DeleteSalesPerson)

}
