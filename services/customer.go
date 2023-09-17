package services

import (
	"errors"
	"log"

	"github.com/NithishNithi/GoShop/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerResponse, error) {
	// Check if a customer with the same customerId or email already exists
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"customerid", user.CustomerId}},
			bson.D{{"email", user.Email}},
		}},
	}

	existingCustomer := &models.Customer{}
	err := p.CustomerCollection.FindOne(p.ctx, filter).Decode(existingCustomer)
	if err == nil {
		return nil, errors.New("Customer with the same customerId or email already exists")
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Insert the new customer
	result, err := p.CustomerCollection.InsertOne(p.ctx, user)
	if err != nil {
		log.Fatal(err)
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
