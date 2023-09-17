package interfaces

import "github.com/NithishNithi/GoShop/models"


type Customer interface {
	CreateCustomer(customer *models.Customer)(*models.CustomerResponse,error)
}