package main

import (
	"context"
	"fmt"
	"net"

	"github.com/NithishNithi/GoShop/constants"
	"github.com/NithishNithi/GoShop/controllers"
	"github.com/NithishNithi/GoShop/database"
	pro "github.com/NithishNithi/GoShop/proto"
	"github.com/NithishNithi/GoShop/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := database.GetCollection(client, "GoShop", "CustomerProfile")
	controllers.CustomerService = services.InitCustomerService(CustomerCollection, context.Background())

}

func main() {
	mongoclient, err := database.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	// grpc server declaration
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	pro.RegisterGoShopServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
