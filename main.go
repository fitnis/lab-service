package main

import (
	"github.com/fitnis/lab-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	lab := router.Group("/lab")
	{
		lab.POST("/collectSample", handlers.CollectSample)
		lab.POST("/recordSample", handlers.RecordSample)
		lab.POST("/processSample", handlers.ProcessSample)
		lab.POST("/evaluateSample", handlers.EvaluateSample)
	}

	router.Run(":8085")
}
