package services

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/NithishNithi/GoTask/constants"
	"github.com/NithishNithi/GoTask/database"
	"github.com/NithishNithi/GoTask/models"

	"github.com/jung-kurt/gofpdf"
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

//  Get task by customerid and save as *pdf ----->
func (p *CustomerService) GetTask(user1 *models.EditTaskDetails) error {
	filter := bson.M{"customerid": user1.CustomerId}
	cursor, err := p.TaskCollection.Find(p.ctx, filter)
	if err != nil {
		return err
	}
	var Tasks []models.Task
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return err
		}
		Tasks = append(Tasks, task)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	for _, task := range Tasks {
		// Add task information to the PDF
		pdf.Cell(0, 10, "Task ID: "+task.TaskId)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Task Title: "+task.Title)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Description: "+task.Description)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Due Date: "+task.DueDate)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Priority: "+task.Priority)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Category: "+task.Category)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Created At: "+task.CreatedAt)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Completed: "+fmt.Sprintf("%v", task.Completed))
		pdf.Ln(10)
		// Add more task fields as needed
		pdf.Ln(10) // Add space between tasks
	}
	// Get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// Construct the full path to the "Downloads" folder
	downloadsFolder := filepath.Join(usr.HomeDir, "Downloads")
	// Ensure the "Downloads" folder exists
	if err := os.MkdirAll(downloadsFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// Generate the PDF filename
	fileName := filepath.Join(downloadsFolder, "tasks.pdf")
	// Create a new file for writing the PDF
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Output the PDF to the file
	err = pdf.Output(file)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// ----------------> CheckTaskDueStatus------------>
func CheckTaskDueStatus() {
	mongoclient, _ := database.ConnectDatabase()
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
			log.Printf("Error while querying tasks: %v\n", err)
			continue
		}
		for cursor.Next(ctx) {
			var task models.Task
			if err := cursor.Decode(&task); err != nil {
				log.Printf("Error decoding task: %v\n", err)
				continue
			}
			task.Completed = true
			update := bson.M{"$set": bson.M{"completed": true}}
			options := options.Update()
			_, err := TaskCollection.UpdateOne(ctx, bson.M{"taskid": task.TaskId}, update, options)
			if err != nil {
				log.Printf("Error updating task: %v\n", err)
			}
			filter := bson.M{
				"customerid": task.CustomerId,
			}
			var customer *models.Customer
			err1 := CustomerCollection.FindOne(ctx, filter).Decode(&customer)
			if err1 != nil {
				return
			}
			go TaskRemainderSMSNotification(task, customer)

		}
		time.Sleep(time.Second)
	}
}

func TaskRemainderSMSNotification(task models.Task, customer *models.Customer) {
	accountSid := constants.AccountSID
	authToken := constants.AuthToken
	to := customer.PhoneNumber
	from := constants.PhoneNumber
	message := "Task ID: " + task.TaskId + ": " + task.Title + ", About: " + task.Description + " has been Completed. This is your Remainder Message from our Team"
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
