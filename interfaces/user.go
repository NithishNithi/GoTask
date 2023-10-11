package interfaces

import "github.com/NithishNithi/GoTask/models"

type Customer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerResponse, error)
	InsertToken(Login *models.Token) (*models.TokenResponse, error)
	CreateTask(Task *models.Task) (*models.Task, error)
	EditTask(Task *models.EditTaskDetails) (*models.Task, error)
	DeleteTask(Task *models.EditTaskDetails)(error)
	GetbyTaskId(Task *models.EditTaskDetails)(*models.Task,error)
	GetTask(Task *models.EditTaskDetails)([]models.Task3, error)
}
