package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/controllers"
	"github.com/NithishNithi/GoTask/database"
	pro "github.com/NithishNithi/GoTask/proto"
	"github.com/NithishNithi/GoTask/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := database.GetCollection(client, "GoTask", "CustomerProfile")
	TokenCollection := database.GetCollection(client, "GoTask", "Tokens")
	TaskCollection := database.GetCollection(client, "GoTask", "TaskManagement")
	controllers.CustomerService = services.InitCustomerService(CustomerCollection, TokenCollection, TaskCollection, context.Background())
}

func main() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	// Configure the log package to write to the log file
	log.SetOutput(logFile)
	// Set a log prefix (optional)
	log.SetPrefix("[YourApp] ")
	// env --->
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}
	dbURL := os.Getenv("DB_URL")
	secretKey := os.Getenv("SECRET_KEY")
	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	phoneNumber := os.Getenv("PHONE_NUMBER")
	constants.ConnectionString = dbURL
	constants.SecretKey = secretKey
	constants.AccountSID = accountSID
	constants.AuthToken = authToken
	constants.PhoneNumber = phoneNumber
	// <----- env
	mongoclient, err := database.ConnectDatabase()
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return
	}
	defer mongoclient.Disconnect(context.TODO())
	initDatabase(mongoclient)
	// grpc server declaration
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	pro.RegisterGoTaskServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Port)

	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
		return
	}
}
