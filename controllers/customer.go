package controllers

import (
	"log"

	"github.com/NithishNithi/GoTask/models"
)

func CreateCustomer(req *models.Customer) (*models.CreateCustomerResponse, error) {
	_, err := CustomerService.CreateCustomer(req)
	if err != nil {
		log.Printf("Error creating customer: %v", err)
		return nil, err
	}

	responseCustomer := models.CreateCustomerResponse{
		Message: "Welcome to GoTask Scheduler, Your Account has Been Created",
	}
	return &responseCustomer, nil
}

func InsertToken(customerid, email, token string) (*models.TokenResponse, error) {
	dbToken := models.Token{Email: email, Token: token, CustomerId: customerid}
	result, err := CustomerService.InsertToken(&dbToken)
	if err != nil {
		log.Printf("Error inserting token: %v", err)
		return nil, err
	}
	return result, nil
}
