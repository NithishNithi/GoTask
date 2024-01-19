package services

import (
	"context"
	"errors"
	"log"

	"github.com/NithishNithi/GoTask/database"
	"github.com/NithishNithi/GoTask/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerResponse, error) {
	user.CustomerId = GenerateUniqueCustomerID()

	// Check if a customer with the same customerId or email already exists
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"customerid", user.CustomerId}},
			bson.D{{"email", user.Email}},
		}},
	}
	existingCustomer := &models.Customer{}
	err := p.CustomerCollection.FindOne(p.ctx, filter).Decode(existingCustomer)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.New("An error occurred while checking for existing customer")
	} else if existingCustomer != nil {
		return nil, errors.New("A customer with the same customerId or email already exists")
	}

	// Insert the new customer
	result, err := p.CustomerCollection.InsertOne(p.ctx, user)
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return nil, errors.New("Failed to create customer")
	}

	// Find and return the newly inserted customer
	newCustomer := &models.Customer{}
	err = p.CustomerCollection.FindOne(p.ctx, bson.M{"_id": result.InsertedID}).Decode(newCustomer)
	if err != nil {
		return nil, err
	}

	response := &models.CustomerResponse{
		CustomerId: newCustomer.CustomerId,
	}

	return response, nil
}

func IsValidUser(user *models.Login) (bool, models.Customer) {
	mongoclient, err := database.ConnectDatabase()
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return false, models.Customer{}
	}

	collection := mongoclient.Database("GoTask").Collection("CustomerProfile")
	query := bson.M{"email": user.Email}
	var customer models.Customer
	err = collection.FindOne(context.TODO(), query).Decode(&customer)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return false, customer
	}
	if customer.Password != user.Password {
		log.Println("Invalid password")
		return false, customer
	}
	return true, customer
}

func (p *CustomerService) InsertToken(user *models.Token) (*models.TokenResponse, error) {
	result, err := p.TokenCollection.InsertOne(p.ctx, &user)
	if err != nil {
		log.Printf("Failed to insert token: %v", err)
		return nil, errors.New("Failed to insert token")
	}

	var newUser models.Token
	query := bson.M{"_id": result.InsertedID}
	err = p.TokenCollection.FindOne(p.ctx, query).Decode(&newUser)
	if err != nil {
		log.Printf("Error fetching inserted token: %v", err)
		return nil, errors.New("Failed to fetch inserted token")
	}

	response := &models.TokenResponse{
		Token: newUser.Token,
	}
	return response, nil
}
