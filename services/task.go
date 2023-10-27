package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/database"
	"github.com/NithishNithi/GoTask/models"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *CustomerService) CreateTask(user *models.Task) (*models.Task, error) {
	_, err := time.Parse("2006-01-02 15:04:05", user.DueDate)
	if err != nil {
		return nil, fmt.Errorf("due date format is invalid: %v", err)
	}
	user.CreatedAt = time.Now().Format(time.RFC850)
	result, err := p.TaskCollection.InsertOne(p.ctx, user)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return nil, errors.New("Failed to create task")
	}
	var newTask *models.Task
	err1 := p.TaskCollection.FindOne(p.ctx, bson.M{"_id": result.InsertedID}).Decode(&newTask)
	if err1 != nil {
		log.Printf("Error fetching created task: %v", err1)
		return nil, errors.New("Failed to fetch created task")
	}
	return newTask, nil
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
		log.Printf("Error while fetching task: %v", err)
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
		log.Printf("Error while updating task: %v", err)
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}
	err1 := p.TaskCollection.FindOne(p.ctx, filter).Decode(&existingTask)
	if err1 != nil {
		log.Printf("Error while fetching task: %v", err1)
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
		log.Printf("Error deleting task: %v", err)
		return err
	}
	if result.DeletedCount == 0 {
		log.Println("Task not deleted")
		return errors.New("Task not deleted")
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
		log.Printf("Error fetching task by ID: %v", err)
		return nil, err
	}
	return result, nil
}

func (p *CustomerService) GetTask(user1 *models.EditTaskDetails) ([]models.Task3, error) {
	filter := bson.M{"customerid": user1.CustomerId}
	cursor, err := p.TaskCollection.Find(p.ctx, filter)
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return nil, err
	}
	var tasks []models.Task3
	for cursor.Next(context.TODO()) {
		var task models.Task3
		err := cursor.Decode(&task)
		if err != nil {
			log.Printf("Error decoding task: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// CheckTaskDueStatus periodically checks for tasks that are past their due date and marks them as completed.
func CheckTaskDueStatus() {
	mongoclient, err := database.ConnectDatabase()
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return
	}

	TaskCollection := mongoclient.Database("GoTask").Collection("TaskManagement")
	CustomerCollection := mongoclient.Database("GoTask").Collection("CustomerProfile")

	for {
		currentTime := time.Now()
		currentTimeStr := currentTime.Format("2006-01-02 15:04:05")

		filter := bson.M{
			"$and": []bson.M{
				{"duedate": bson.M{"$lt": currentTimeStr}},
				{"completed": false},
			},
		}

		ctx := context.TODO()
		cursor, err := TaskCollection.Find(ctx, filter)
		if err != nil {
			log.Printf("Error while querying tasks: %v", err)
			continue
		}

		for cursor.Next(ctx) {
			var task models.Task
			if err := cursor.Decode(&task); err != nil {
				log.Printf("Error decoding task: %v", err)
				continue
			}

			task.Completed = true
			update := bson.M{"$set": bson.M{"completed": true}}
			options := options.Update()
			_, err := TaskCollection.UpdateOne(ctx, bson.M{"taskid": task.TaskId}, update, options)
			if err != nil {
				log.Printf("Error updating task: %v", err)
			}

			filter := bson.M{
				"customerid": task.CustomerId,
			}

			var customer *models.Customer
			err1 := CustomerCollection.FindOne(ctx, filter).Decode(&customer)
			if err1 != nil {
				log.Printf("Error fetching customer: %v", err1)
				return
			}

			go TaskReminderSMSNotification(task, customer)
		}

		time.Sleep(time.Second)
	}
}

func TaskReminderSMSNotification(task models.Task, customer *models.Customer) {
	accountSid := constants.AccountSID
	authToken := constants.AuthToken
	to := customer.PhoneNumber
	from := constants.PhoneNumber
	message := "Task ID: " + task.TaskId + ": " + task.Title + ", About: " + task.Description + " has been Completed. This is your Reminder Message from our Team"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Retry up to 3 times with a delay of 5 seconds between retries
	for attempt := 1; attempt <= 3; attempt++ {
		params := &twilioApi.CreateMessageParams{
			To:   &to,
			From: &from,
			Body: &message,
		}

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			log.Printf("Attempt %d: Error sending SMS message: %v", attempt, err)
			if isTimeoutError(err) {
				// Sleep for 5 seconds before retrying
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
		break
	}
}

func isTimeoutError(err error) bool {
	netErr, isNetErr := err.(net.Error)
	return isNetErr && netErr.Timeout()
}
