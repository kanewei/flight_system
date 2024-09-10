package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AirportHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type airportHandler struct {
	airportService *service.AirportService
}

func NewAirportHandler() AirportHandler {
	return &airportHandler{
		airportService: service.NewAirportService(),
	}
}

func (a *airportHandler) Create(ctx *gin.Context) {
	req := new(model.CreateAirportRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := a.airportService.CreateAirport(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.Status(http.StatusOK)
}

func (a *airportHandler) Get(ctx *gin.Context) {
	req := new(model.GetAirportRequest)
	if err := ctx.BindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Code != "" {
		airport, err := a.airportService.GetAirportByCode(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"airport": airport})
	}

	airports, err := a.airportService.GetAirportByCity(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"airports": airports})
}
