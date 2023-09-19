package services

import (
	"time"

	"github.com/NithishNithi/GoTask/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *CustomerService) CreateTask(user *models.Task) (*models.Task, error) {
	user.CreatedAt = time.Now().Format(time.RFC850)
	result, err := p.TaskCollection.InsertOne(p.ctx, user)
	if err != nil {
		return nil, err
	}
	var newtask *models.Task
	err1 := p.TaskCollection.FindOne(p.ctx, bson.M{"_id": result.InsertedID}).Decode(&newtask)
	if err1 != nil {
		return nil, err1
	}

	return newtask, nil

}
