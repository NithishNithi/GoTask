package services

import (
	"fmt"
	"time"

	"github.com/NithishNithi/GoTask/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (p *CustomerService) EditTask(user *models.EditTaskDetails) (*models.Task, error) {
	filter := bson.M{
		"customerid": user.CustomerId,
		"taskid":     user.TaskId,
	}
	update := bson.M{"$set": bson.M{user.Field: user.Value}}
	var existingTask *models.Task
	err := p.TaskCollection.FindOne(p.ctx, filter).Decode(&existingTask)
	if err != nil {
		fmt.Println("error while fetching task")
		return nil, err
	}
	newUpdate := models.Update{
		UpdatedAt: time.Now().Format(time.RFC850),
		Changes:   fmt.Sprintf("Field '%s' updated to '%s'", user.Field, user.Value),
	}
	existingTask.UpdateHistory = append(existingTask.UpdateHistory, newUpdate)
	update["$set"].(bson.M)["updatehistory"] = existingTask.UpdateHistory
	options := options.Update()
	result, err := p.TaskCollection.UpdateOne(p.ctx, filter, update, options)
	if err != nil {
		fmt.Println("error while updating")
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}
	err1 := p.TaskCollection.FindOne(p.ctx, filter).Decode(&existingTask)
	if err1 != nil {
		fmt.Println("error while fetching task")
		return nil, err1
	}
	return existingTask, nil
}

func (p *CustomerService) DeleteTask(user *models.EditTaskDetails) error {
	filter := bson.M{
		"customerid": user.CustomerId,
		"taskid":     user.TaskId,
	}
	result, err := p.TaskCollection.DeleteOne(p.ctx, filter)
	if err != nil {
		return fmt.Errorf("error: Task not deleted")
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("error: Task not deleted")
	}
	return nil
}

func (p *CustomerService) GetbyTaskId(user *models.EditTaskDetails) (*models.Task, error) {
	filter := bson.M{
		"customerid": user.CustomerId,
		"taskid":     user.TaskId,
	}
	var result *models.Task
	err := p.TaskCollection.FindOne(p.ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
