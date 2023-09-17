package controllers

import (
	"github.com/NithishNithi/GoShop/interfaces"
	"github.com/gin-gonic/gin"
	pro "github.com/NithishNithi/GoShop/proto"
)

type RPCServer struct {
	pro.UnimplementedGoShopServiceServer
}

var (
	ctx             gin.Context
	CustomerService interfaces.Customer
)
