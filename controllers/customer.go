package controllers

import (
	"context"

	"github.com/NithishNithi/GoTask/models"
	pro "github.com/NithishNithi/GoTask/proto"
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	dbCustomer := models.Customer{
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		DateofBirth: req.DateofBirth,
		PhoneNumber: req.PhoneNumber,
		HouseNo:     req.HouseNo,
		Street:      req.Street,
		City:        req.City,
		Country:     req.Country,
		Zip:         req.Zip,
	}

	result, err := CustomerService.CreateCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
			Success:    true,
			Message:    "Welcome to GoTask Scheduler, Your Account has Been Created",
		}
		return responseCustomer, nil
	}
}

func (s *RPCServer) InsertToken(ctx context.Context, req *pro.Token) (*pro.TokenResponse, error) {
	dbCustomer := models.Token{Email: req.Email, Token: req.Token, CustomerId: req.CustomerId}
	result, err := CustomerService.InsertToken(&dbCustomer)
	if err != nil {
		return nil, err
	} else {

		responsetoken := &pro.TokenResponse{
			Token: result.Token,
		}
		return responsetoken, nil
	}
}
