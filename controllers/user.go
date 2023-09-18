package controllers

import (
	"github.com/NithishNithi/GoTask/interfaces"
	pro "github.com/NithishNithi/GoTask/proto"
	"github.com/gin-gonic/gin"
)

type RPCServer struct {
	pro.UnimplementedGoTaskServiceServer
}

var (
	ctx             gin.Context
	CustomerService interfaces.Customer
)
