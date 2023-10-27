package handlers

import (
	"log"
	"net/http"

	grpcclient "github.com/NithishNithi/GoTask/cmd/grpc"
	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/models"
	pb "github.com/NithishNithi/GoTask/proto"
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

	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := Client.CreateTask(c.Request.Context(), &pb.TaskDetails{TaskId: request.TaskId, CustomerId: Customerid, Title: request.Title, Description: request.Description, DueDate: request.DueDate, Priority: request.Priority, Category: request.Category, CreatedAt: request.CreatedAt, Completed: request.Completed})
	if err != nil {
		log.Printf("Error creating task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": response})
}

func EditTask(c *gin.Context) {
	taskid := c.Param("taskid")
	token := c.GetHeader("Authorization")
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
	var request *models.EditTaskDetails
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := Client.EditTask(c.Request.Context(), &pb.EditTaskDetails{CustomerId: Customerid, TaskId: taskid, Field: request.Field, Value: request.Value})
	if err != nil {
		log.Printf("Error editing task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": response})
}

func DeleteTask(c *gin.Context) {
	taskid := c.Param("taskid")
	token := c.GetHeader("Authorization")
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
	Client, _ := grpcclient.GetGrpcClientInstance()
	_, err = Client.DeleteTask(c.Request.Context(), &pb.TaskDelete{TaskId: taskid, CustomerId: customerid})
	if err != nil {
		log.Printf("Error deleting task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Task Deleted"})
	}
}

func GetbyTaskId(c *gin.Context) {
	taskid := c.Param("taskid")
	token := c.GetHeader("Authorization")
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
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := Client.GetTaskbyId(c.Request.Context(), &pb.TaskDelete{TaskId: taskid, CustomerId: customerid})
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
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err := Client.GetTask(c.Request.Context(), &pb.TaskDelete{CustomerId: customerid})
	if err != nil {
		log.Printf("Error getting task: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Tasks not found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": response})
	}
}
