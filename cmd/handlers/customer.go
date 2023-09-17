package handlers

import (
	"log"
	"net/http"

	grpcclient "github.com/NithishNithi/GoShop/cmd/grpc"
	pb "github.com/NithishNithi/GoShop/proto"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var request pb.CustomerDetails
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Client, _ := grpcclient.GetGrpcClientInstance()
	response, err1 := Client.CreateCustomer(c.Request.Context(), &request)
	if err1 != nil {
		log.Fatal(err1)
	}
	c.JSON(http.StatusOK, gin.H{"value": response})
}
