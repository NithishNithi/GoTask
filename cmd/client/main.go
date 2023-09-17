package main

import (
	"fmt"

	grpcclient "github.com/NithishNithi/GoShop/cmd/grpc"
	"github.com/NithishNithi/GoShop/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Client Server is Running")
	_, conn := grpcclient.GetGrpcClientInstance()
	defer conn.Close()
	r := gin.Default()
	routes.SetUpRoutes(r)
	r.Run(":8000")

}
