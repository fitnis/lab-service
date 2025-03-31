package handlers

import (
	"net/http"

	"github.com/fitnis/lab-service/models"
	"github.com/fitnis/lab-service/services"
	"github.com/gin-gonic/gin"
)

// curl -X POST -H "Content-Type: application/json" -d '{"sampleId":"abc123","patientId":"123","type":"blood"}' http://localhost:8085/lab/collectSample
func CollectSample(c *gin.Context) {
	var req models.Sample
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sample"})
		return
	}
	c.JSON(http.StatusCreated, services.CollectSample(req))
}

// curl -X POST -H "Content-Type: application/json" -d '{"sampleId":"abc123","patientId":"123","type":"blood"}' http://localhost:8085/lab/recordSample
func RecordSample(c *gin.Context) {
	var req models.Sample
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sample"})
		return
	}
	c.JSON(http.StatusCreated, services.RecordSample(req))
}

// curl -X POST -H "Content-Type: application/json" -d '{"sampleId":"abc123","patientId":"123","type":"blood"}' http://localhost:8085/lab/processSample
func ProcessSample(c *gin.Context) {
	var req models.Sample
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sample"})
		return
	}
	c.JSON(http.StatusCreated, services.ProcessSample(req))
}

// curl -X POST -H "Content-Type: application/json" -d '{"sampleId":"abc123","result":"Normal"}' http://localhost:8085/lab/evaluateSample
func EvaluateSample(c *gin.Context) {
	var req models.SampleEvaluation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid evaluation"})
		return
	}
	c.JSON(http.StatusCreated, services.EvaluateSample(req))
}
