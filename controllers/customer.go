package controllers

import (
	"context"

	"github.com/NithishNithi/GoShop/models"
	pro "github.com/NithishNithi/GoShop/proto"
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
		DateofBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Address: []*models.Address{&address},
	}

	result,err:=CustomerService.CreateCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
			Success: true,
			Message: "Welcome to GoShop Your Account has Been Created",
		}
		return responseCustomer, nil
	}
}
