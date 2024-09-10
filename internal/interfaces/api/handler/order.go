package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	CreateOrder(ctx *gin.Context)
}

type orderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler() OrderHandler {
	return &orderHandler{
		orderService: service.NewOrderService(),
	}
}

func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	req := new(model.CreateOrderRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := o.orderService.CreateOrder(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
