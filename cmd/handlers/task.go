package handlers

import (
	"net/http"

	grpcclient "github.com/NithishNithi/GoTask/cmd/grpc"
	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/models"
	pb "github.com/NithishNithi/GoTask/proto"
	"github.com/NithishNithi/GoTask/services"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request *models.Task
	err1 := c.ShouldBindJSON(&request)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err2 := Client.CreateTask(c.Request.Context(), &pb.TaskDetails{TaskId: request.TaskId, CustomerId: Customerid, Title: request.Title, Description: request.Description, DueDate: request.DueDate, Priority: request.Priority, Category: request.Category, CreatedAt: request.CreatedAt, Completed: request.Completed})
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request *models.EditTaskDetails
	err1 := c.ShouldBindJSON(&request)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err2 := Client.EditTask(c.Request.Context(), &pb.EditTaskDetails{CustomerId: Customerid, TaskId: taskid, Field: request.Field, Value: request.Value})
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	_, err2 := Client.DeleteTask(c.Request.Context(), &pb.TaskDelete{TaskId: taskid, CustomerId: customerid})
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err2 := Client.GetTaskbyId(c.Request.Context(), &pb.TaskDelete{TaskId: taskid, CustomerId: customerid})
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err2})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": response})

}

func GetTask(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	customerid, err := services.ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	_, err1 := Client.GetTask(c.Request.Context(), &pb.TaskDelete{CustomerId: customerid})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err1})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Tasks Found"})
	}

}
