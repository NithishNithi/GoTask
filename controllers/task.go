package controllers

import (
	"log"

	"github.com/NithishNithi/GoTask/models"
	pro "github.com/NithishNithi/GoTask/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateTask(req models.Task1) (*pro.TaskResponse, error) {

	dbtask := models.Task{
		TaskId:      req.TaskId,
		CustomerId:  req.CustomerId,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Priority:    req.Priority,
		Category:    req.Category,
		CreatedAt:   req.CreatedAt,
		Completed:   req.Completed,
	}
	result, err := CustomerService.CreateTask(&dbtask)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, status.Error(codes.Internal, "Failed to create task")
	}
	responsetask := &pro.TaskResponse{
		TaskId:  result.TaskId,
		Title:   result.Title,
		DueDate: result.DueDate,
	}
	return responsetask, nil
}

func EditTask(req *models.EditTaskDetails) error {

	dbtask := models.EditTaskDetails{
		TaskId:     req.TaskId,
		CustomerId: req.CustomerId,
		Field:      req.Field,
		Value:      req.Value,
	}
	_, err := CustomerService.EditTask(&dbtask)
	if err != nil {
		log.Printf("Error editing task: %v", err)
		return status.Error(codes.Internal, "Failed to edit task")
	}

	return nil
}

func DeleteTask(taskid, customerid string) error {
	dbdelete := models.EditTaskDetails{
		TaskId:     taskid,
		CustomerId: customerid,
	}
	err := CustomerService.DeleteTask(&dbdelete)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return status.Error(codes.Internal, "Failed to delete task")
	}
	return nil
}

func GetTaskbyId(taskid, customerid string) (*models.Task, error) {
	dbgettaskbyid := models.EditTaskDetails{
		TaskId:     taskid,
		CustomerId: customerid,
	}
	result, err := CustomerService.GetbyTaskId(&dbgettaskbyid)
	if err != nil {
		log.Printf("Error getting task by ID: %v", err)
		return nil, status.Error(codes.NotFound, "Task not found")
	}
	return result, nil
}

func GetTask(customerid string) ([]models.Task3, error) {
	if customerid == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid request")
	}
	dbgettask := models.EditTaskDetails{
		CustomerId: customerid,
	}
	tasks, err := CustomerService.GetTask(&dbgettask)
	if err != nil {
		log.Printf("Error getting tasks: %v", err)
		return nil, status.Error(codes.Internal, "Failed to get tasks")
	}

	return tasks, nil
}
