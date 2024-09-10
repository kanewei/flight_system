package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightHandler interface {
	Create(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Search(ctx *gin.Context)
}

type flightHandler struct {
	flightService *service.FlightService
}

func NewFlightHandler() FlightHandler {
	return &flightHandler{
		flightService: service.NewFlightService(),
	}
}

func (f *flightHandler) Create(ctx *gin.Context) {
	req := new(model.CreateFlightRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := f.flightService.CreateFlight(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (f *flightHandler) GetById(ctx *gin.Context) {
	req := new(model.GetFlightRequest)
	if err := ctx.BindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := f.flightService.GetFlightById(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (f *flightHandler) Search(ctx *gin.Context) {
	req := new(model.SearchFlightRequest)
	if err := ctx.BindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	flights, err := f.flightService.SearchFlight(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"flights": flights})
}
