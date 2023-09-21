package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/NithishNithi/GoTask/database"
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

// ----------------> CheckTaskDueStatus------------>

func CheckTaskDueStatus() {
	mongoclient, _ := database.ConnectDatabase()
	TaskCollection := mongoclient.Database("GoTask").Collection("TaskManagement")
	for {
		// Parse the current time in the same format as your due date
		currentTime := time.Now()
		currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
		currentTimeUTC, err := time.Parse("2006-01-02 15:04:05", currentTimeStr)
		fmt.Println(currentTimeUTC)
		if err != nil {
			log.Printf("Error parsing current time: %v\n", err)
			continue
		}

		// Calculate the start of the next minute
		nextMinute := currentTimeUTC.Add(time.Minute)
		nextMinute = time.Date(nextMinute.Year(), nextMinute.Month(), nextMinute.Day(), nextMinute.Hour(), nextMinute.Minute(), 0, 0, nextMinute.Location())

		// Calculate the duration until the start of the next minute
		sleepDuration := nextMinute.Sub(currentTimeUTC)

		// Sleep until the start of the next minute
		time.Sleep(sleepDuration)

		// Fetch tasks where DueTime has passed and the task is not already marked as completed
		filter := bson.M{
			"$and": []bson.M{
				{"duedate": bson.M{"$lt": currentTimeStr}},
				{"completed": false},
			},
		}
		ctx := context.TODO()
		cursor, err := TaskCollection.Find(ctx, filter)
		fmt.Println("cursor:",cursor)
		if err != nil {
			log.Printf("Error while querying tasks: %v\n", err)
			continue
		}

		for cursor.Next(ctx) {
			var task models.Task
			if err := cursor.Decode(&task); err != nil {
				log.Printf("Error decoding task: %v\n", err)
				continue
			}
			// Mark the task as completed
			task.Completed = true

			// Update the task's completion status in the database
			update := bson.M{"$set": bson.M{"completed": true}}
			options := options.Update()
			response, err := TaskCollection.UpdateOne(ctx, bson.M{"taskid": task.TaskId }, update, options)
			fmt.Println("cus:",task.CustomerId," task:",task.TaskId)
			fmt.Println("1",response.ModifiedCount)
			fmt.Println("2",response.MatchedCount)
			fmt.Println("3",response.UpsertedCount)
			if err != nil {
				log.Printf("Error updating task: %v\n", err)
			}
		}
	}
}

func RunTaskDueStatusChecker() {
	fmt.Println("1")
	go CheckTaskDueStatus()
}

