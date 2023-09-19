package controllers

import (
	"context"

	"github.com/NithishNithi/GoTask/models"
	pro "github.com/NithishNithi/GoTask/proto"
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
		UpdatedAt:   req.UpdatedAt,
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
