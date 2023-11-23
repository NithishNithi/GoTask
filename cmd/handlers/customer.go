package handlers

import (
	"log"
	"net/http"
	"fmt"
	"github.com/NithishNithi/GoTask/controllers"
	"github.com/NithishNithi/GoTask/models"
	"github.com/NithishNithi/GoTask/services"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var request models.Customer
	if err := c.ShouldBindJSON(&request); err != nil {
		// Log the error but continue processing
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	// Client, _ := grpcclient.GetGrpcClientInstance()
	response, err1 := controllers.CreateCustomer(&request)
	fmt.Print(response)
	if err1 != nil {
		// Log the error but continue processing
		log.Println("Error creating customer:", err1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": response})
}

func LoginCustomer(c *gin.Context) {
	var request *models.Login
	err := c.ShouldBindJSON(&request)
	if err != nil {
		// Log the error but continue processing
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	status, ans := services.IsValidUser(request)
	if status {
		token, err := services.CreateToken(request.Email, ans.CustomerId)
		if err != nil {
			// Log the error but continue processing
			log.Println("Error creating token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		// Client, _ := grpcclient.GetGrpcClientInstance()
		response, err1 := controllers.InsertToken(request.CustomerId, request.Email, token)
		if err1 != nil {
			// Log the error but continue processing
			log.Println("Error inserting token:", err1)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": response.Token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
