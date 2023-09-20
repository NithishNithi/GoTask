package models

type Task struct {
	TaskId        string   `json:"taskid" bson:"taskid"`
	CustomerId    string   `json:"customerid" bson:"customerid"`
	Title         string   `json:"title" bson:"title"`
	Description   string   `json:"description" bson:"description"`
	DueDate       string   `json:"duedate" bson:"duedate"`
	Priority      string   `json:"priority" bson:"priority"`
	Category      string   `json:"category" bson:"category"`
	CreatedAt     string   `json:"createdat" bson:"createdat"`
	Completed     bool     `json:"completed" bson:"completed"`
	UpdateHistory []Update `json:"updatehistory" bson:"updatehistory"`
}

type Update struct {
	UpdatedAt string `json:"updatedat"`
	Changes   string `json:"changes"`
}

type EditTaskDetails struct {
	TaskId     string `json:"taskid" bson:"taskid"`
	CustomerId string `json:"customerid" bson:"customerid"`
	Field      string `json:"field"`
	Value      string `json:"value"`
}
