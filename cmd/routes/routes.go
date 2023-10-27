package routes

import (
	"github.com/NithishNithi/GoTask/cmd/handlers"
	"github.com/NithishNithi/GoTask/services"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	services.RunTaskDueStatusChecker()

	r.Static("/signup", "/home/nithish/go/src/GoTask/frontend/signup")
	r.Static("/signin", "/home/nithish/go/src/GoTask/frontend/signin")
	r.Static("/home", "/home/nithish/go/src/GoTask/frontend/home")

	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.CreateCustomer)
		auth.POST("/login", handlers.LoginCustomer)
	}

	task := r.Group("/tasks")
	{
		task.POST("/createtask", handlers.CreateTask)
		task.POST("/edittask/:taskid", handlers.EditTask)
		task.GET("/deletetask/:taskid", handlers.DeleteTask)
		task.GET("/gettaskbyid/:taskid", handlers.GetbyTaskId)
		task.POST("/gettask", handlers.GetTask)
	}
}
