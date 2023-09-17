package services

import (
	"context"

	"github.com/NithishNithi/GoShop/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.Customer {
	return &CustomerService{collection, ctx}
}
