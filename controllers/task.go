package controllers

import (
	"context"
	"log"

	"github.com/NithishNithi/GoTask/models"
	pro "github.com/NithishNithi/GoTask/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *RPCServer) CreateTask(ctx context.Context, req *pro.TaskDetails) (*pro.TaskResponse, error) {
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

func (s *RPCServer) EditTask(ctx context.Context, req *pro.EditTaskDetails) (*pro.TaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Request is nil")
	}
	if req.CustomerId == "" || req.Field == "" || req.TaskId == "" || req.Value == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing required fields")
	}

	dbtask := models.EditTaskDetails{
		TaskId:     req.TaskId,
		CustomerId: req.CustomerId,
		Field:      req.Field,
		Value:      req.Value,
	}
	result, err := CustomerService.EditTask(&dbtask)
	if err != nil {
		log.Printf("Error editing task: %v", err)
		return nil, status.Error(codes.Internal, "Failed to edit task")
	}
	responsetask := &pro.TaskResponse{
		TaskId:  result.TaskId,
		Title:   result.Title,
		DueDate: result.DueDate,
	}
	return responsetask, nil
}

func (s *RPCServer) DeleteTask(ctx context.Context, req *pro.TaskDelete) (*pro.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid request")
	}
	dbdelete := models.EditTaskDetails{
		TaskId:     req.TaskId,
		CustomerId: req.CustomerId,
	}
	err := CustomerService.DeleteTask(&dbdelete)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return nil, status.Error(codes.Internal, "Failed to delete task")
	}
	return &pro.Empty{}, nil
}

func (s *RPCServer) GetTaskbyId(ctx context.Context, req *pro.TaskDelete) (*pro.TaskDetails, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid request")
	}
	dbgettaskbyid := models.EditTaskDetails{
		TaskId:     req.TaskId,
		CustomerId: req.CustomerId,
	}
	result, err := CustomerService.GetbyTaskId(&dbgettaskbyid)
	if err != nil {
		log.Printf("Error getting task by ID: %v", err)
		return nil, status.Error(codes.NotFound, "Task not found")
	}
	responsetask := &pro.TaskDetails{
		TaskId:      result.TaskId,
		CustomerId:  result.CustomerId,
		Title:       result.Title,
		Description: result.Description,
		DueDate:     result.DueDate,
		Priority:    result.Priority,
		Category:    result.Category,
		CreatedAt:   result.CreatedAt,
		Completed:   result.Completed,
	}
	return responsetask, nil
}

func (s *RPCServer) GetTask(ctx context.Context, req *pro.TaskDelete) (*pro.GetTasksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid request")
	}
	dbgettask := models.EditTaskDetails{
		CustomerId: req.CustomerId,
	}
	tasks, err := CustomerService.GetTask(&dbgettask)
	if err != nil {
		log.Printf("Error getting tasks: %v", err)
		return nil, status.Error(codes.Internal, "Failed to get tasks")
	}
	var taskDetails []*pro.TaskDetails
	for _, task := range tasks {
		taskDetail := &pro.TaskDetails{
			TaskId:      task.TaskId,
			CustomerId:  task.CustomerId,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Priority:    task.Priority,
			Category:    task.Category,
			CreatedAt:   task.CreatedAt,
			Completed:   task.Completed,
		}
		taskDetails = append(taskDetails, taskDetail)
	}

	response := &pro.GetTasksResponse{
		Tasks: taskDetails,
	}
	return response, nil
}
