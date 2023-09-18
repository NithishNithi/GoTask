package services

import (
	"context"

	"github.com/NithishNithi/GoTask/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	TokenCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection1 *mongo.Collection,collection2 *mongo.Collection, ctx context.Context) interfaces.Customer {
	return &CustomerService{collection1,collection2, ctx}
}
