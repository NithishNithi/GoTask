package handlers

import (
	"log"
	"net/http"

	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/controllers"
	"github.com/NithishNithi/GoTask/models"
	"github.com/NithishNithi/GoTask/services"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var request *models.Task1
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	request.CustomerId = Customerid
	// Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := controllers.CreateTask(*request)
	if err != nil {
		log.Printf("Error creating task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": response})
}

func EditTask(c *gin.Context) {
	var request *models.EditTask
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}

	customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	task := &models.EditTaskDetails{
		TaskId:     request.TaskId,
		CustomerId: customerid,
		Field:      request.Field,
		Value:      request.Value,
	}

	// Client, _ := grpcclient.GetGrpcClientInstance()
	err = controllers.EditTask(task)
	if err != nil {
		log.Printf("Error editing task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func DeleteTask(c *gin.Context) {
	var request *models.Task1
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	// Client, _ := grpcclient.GetGrpcClientInstance()
	err = controllers.DeleteTask(request.TaskId, customerid)
	if err != nil {
		log.Printf("Error deleting task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Task Deleted"})
	}
}

func GetbyTaskId(c *gin.Context) {
	var request *models.Task1
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	// Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := controllers.GetTaskbyId(request.TaskId, customerid)
	if err != nil {
		log.Printf("Error getting task by ID: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": response})
}

func GetTask(c *gin.Context) {
	var request *models.Task1
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	// Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := controllers.GetTask(customerid)
	if err != nil {
		log.Printf("Error getting task: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Tasks not found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": response})
	}
}
