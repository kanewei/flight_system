package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AirplaneHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type airplaneHandler struct {
	airplaneService *service.AirplaneService
}

func NewAirplaneHandler() AirplaneHandler {
	return &airplaneHandler{
		airplaneService: service.NewAirplaneService(),
	}
}

func (a *airplaneHandler) Create(ctx *gin.Context) {
	req := new(model.CreateAirplaneRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := a.airplaneService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (a *airplaneHandler) Get(ctx *gin.Context) {
	req := new(model.GetAirplaneRequest)
	if err := ctx.BindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.ID != 0 {
		airplane, err := a.airplaneService.GetAirplaneById(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"airplane": airplane})
	}

	airplanes, err := a.airplaneService.GetAirplaneByModel(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"airplanes": airplanes})
}
