package main

import (
	"github.com/fitnis/lab-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	lab := router.Group("/lab")
	{
		handlers.RegisterLabRoutes(lab)
	}

	router.Run(":8085")
}
