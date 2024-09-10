package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler interface {
	CreateTicketOrder(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetByUserId(ctx *gin.Context)
}

type ticketHandler struct {
	ticketService *service.TicketService
}

func NewTicketHandler() TicketHandler {
	return &ticketHandler{
		ticketService: service.NewTicketService(),
	}
}

func (t *ticketHandler) CreateTicketOrder(ctx *gin.Context) {
	req := new(model.CreateTicketOrderRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := t.ticketService.CreateTicketOrder(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (t *ticketHandler) GetById(ctx *gin.Context) {
	req := new(model.GetTicketRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ticket, err := t.ticketService.GetById(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

func (t *ticketHandler) GetByUserId(ctx *gin.Context) {
	req := new(model.GetUserTicketRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ticket, err := t.ticketService.GetByUserId(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"ticket": ticket})
}
