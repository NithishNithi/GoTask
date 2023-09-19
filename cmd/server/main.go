package main

import (
	"context"
	"fmt"
	"net"

	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/controllers"
	"github.com/NithishNithi/GoTask/database"
	pro "github.com/NithishNithi/GoTask/proto"
	"github.com/NithishNithi/GoTask/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := database.GetCollection(client, "GoTask", "CustomerProfile")
	TokenCollection:=database.GetCollection(client,"GoTask","Tokens")
	TaskCollection := database.GetCollection(client,"GoTask","TaskManagement")
	controllers.CustomerService = services.InitCustomerService(CustomerCollection,TokenCollection,TaskCollection,context.Background())

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
	pro.RegisterGoTaskServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
