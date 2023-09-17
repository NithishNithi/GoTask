package main

import (
	"log"
	"net/http"

	pb "github.com/NithishNithi/GoShop/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// var (
// 	mongoclient *mongo.Client
// 	ctx         context.Context
// 	server      *gin.Engine
// )

func main() {
	r := gin.Default()
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	Client := pb.NewGoShopServiceClient(conn)
	r.POST("signup", func(c *gin.Context) {
		var request pb.CustomerDetails

		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err1 := Client.CreateCustomer(c.Request.Context(), &request)
		if err1 != nil {
			log.Fatal(err1)
		}
		c.JSON(http.StatusOK, gin.H{"value": response})
	})
	r.Run(":8000")
}
