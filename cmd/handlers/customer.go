package handlers

import (
	"fmt"
	"log"
	"net/http"

	grpcclient "github.com/NithishNithi/GoTask/cmd/grpc"
	"github.com/NithishNithi/GoTask/models"
	pb "github.com/NithishNithi/GoTask/proto"
	"github.com/NithishNithi/GoTask/services"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var request pb.CustomerDetails
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err1 := Client.CreateCustomer(c.Request.Context(), &request)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1})
	}
	c.JSON(http.StatusOK, gin.H{"value": response})
}

func LoginCustomer(c *gin.Context) {
	var request *models.Login
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, ans := services.IsValidUser(request)
	if status {
		token, err := services.CreateToken(request.Email, ans.CustomerId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
			return
		}
		Client, _ := grpcclient.GetGrpcClientInstance()
		response, err1 := Client.InsertToken(c.Request.Context(), &pb.Token{CustomerId: request.CustomerId, Email: request.Email, Token: token})
		if err1 != nil {
			log.Fatal(err1)
		}
		c.JSON(http.StatusOK, gin.H{"token": response.Token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
