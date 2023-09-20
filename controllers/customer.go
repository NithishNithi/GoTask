package controllers

import (
	"context"

	"github.com/NithishNithi/GoTask/models"
	pro "github.com/NithishNithi/GoTask/proto"
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	var address models.Address
	if req != nil {
		address = models.Address{
			Country: req.Address[0].Country,
			Street:  req.Address[0].Street,
			City:    req.Address[0].City,
			State:   req.Address[0].State,
			Zip:     req.Address[0].Zip,
		}
	}
	dbCustomer:=models.Customer{
		CustomerId: req.CustomerId,
		FullName: req.FullName,
		Email: req.Email,
		Password: req.Password,
		DateofBirth: req.DateofBirth,
		PhoneNumber: req.PhoneNumber,
		Address: []models.Address{address},
	}

	result,err:=CustomerService.CreateCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
			Success: true,
			Message: "Welcome to GoTask Scheduler, Your Account has Been Created",
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

		responsetoken:=&pro.TokenResponse{
			Token: result.Token,
		}
		return responsetoken,nil
	}
}
