package controllers

import (
	"context"
	"fmt"

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
		return nil, err
	} else {
		responsetask := &pro.TaskResponse{
			TaskId:  result.TaskId,
			Title:   result.Title,
			DueDate: result.TaskId,
		}
		return responsetask, nil
	}

}

func (s *RPCServer) EditTask(ctx context.Context, req *pro.EditTaskDetails) (*pro.TaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Request is nil")
	}
	if req.CustomerId == "" || req.Field == "" || req.TaskId == "" || req.Value == "" {
		fmt.Println("error fpound")
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
		return nil, err
	} else {
		responsetask := &pro.TaskResponse{
			TaskId:  result.TaskId,
			Title:   result.Title,
			DueDate: result.TaskId,
		}
		return responsetask, nil

	}
}
