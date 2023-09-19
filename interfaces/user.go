package interfaces

import "github.com/NithishNithi/GoTask/models"

type Customer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerResponse, error)
	InsertToken(Login *models.Token) (*models.TokenResponse, error)
	CreateTask(Task *models.Task) (*models.Task, error)
}
