package models

import "time"

type Task struct {
	TaskId        string   `json:"taskid" bson:"taskid"`
	CustomerId    string   `json:"customerid" bson:"customerid"`
	Title         string   `json:"title" bson:"title"`
	Description   string   `json:"description" bson:"description"`
	DueDate       string   `json:"duedate" bson:"duedate"`
	Priority      string   `json:"priority" bson:"priority"`
	Category      string   `json:"category" bson:"category"`
	CreatedAt     string   `json:"createdat" bson:"createdat"`
	UpdatedAt     string   `json:"updatedat" bson:"updatedat"`
	Completed     bool     `json:"completed" bson:"completed"`
	UpdateHistroy []Update `json:"updatehistroy" bson:"updatehistroy"`
}

type Update struct {
	UpdatedAt time.Time `json:"updated_at"`
	Changes   string    `json:"changes"`
}
