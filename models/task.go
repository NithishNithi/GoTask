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

// it is for getting token from frontend
type Task1 struct {
	Token         string   `json:"token" bson:"token"`
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

type Task3 struct {
	TaskId      string `json:"taskid" bson:"taskid"`
	CustomerId  string `json:"customerid" bson:"customerid"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	DueDate     string `json:"duedate" bson:"duedate"`
	Priority    string `json:"priority" bson:"priority"`
	Category    string `json:"category" bson:"category"`
	CreatedAt   string `json:"createdat" bson:"createdat"`
	Completed   bool   `json:"completed" bson:"completed"`
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

type TaskDetails struct {
}

type EditTask struct {
	Token  string `json:"token" bson:"token"`
	TaskId string `json:"taskid" bson:"taskid"`
	Field  string `json:"field" bson:"field"`
	Value  string `json:"value" bson:"value"`
}
