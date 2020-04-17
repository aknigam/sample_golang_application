package main

import (
	"github.com/gin-gonic/gin"
	"sample_golang_application/routers"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/demo/service/v1/")
	routers.SetApplicationRoutes(v1)
	router.Run()

}
