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
        "taskid":      user.TaskId,
    }
    update := bson.M{"$set": bson.M{user.Field: user.Value}}

    // Fetch the current task to access its existing update history
    var existingTask *models.Task
    err := p.TaskCollection.FindOne(p.ctx, filter).Decode(&existingTask)
    if err != nil {
        fmt.Println("error while fetching task")
        return nil, err
    }

    // Create a new update record
    newUpdate := models.Update{
        UpdatedAt: time.Now().Format(time.RFC850),
        Changes:   fmt.Sprintf("Field '%s' updated to '%s'", user.Field, user.Value),
    }

    // Append the new update record to the existing update history
    existingTask.UpdateHistory = append(existingTask.UpdateHistory, newUpdate)

    // Update the task with the new field value and update history
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

	return existingTask,nil
}
