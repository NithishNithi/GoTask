package routes

import (
	"github.com/NithishNithi/GoShop/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.CreateCustomer)
	}
}
