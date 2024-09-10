package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffHandler interface {
	CreateStaff(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type staffHandler struct {
	staffService *service.StaffService
}

func NewStaffHandler() StaffHandler {
	return &staffHandler{
		staffService: service.NewStaffService(),
	}
}

func (s *staffHandler) CreateStaff(ctx *gin.Context) {
	req := new(model.CreateStaffRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := s.staffService.CreateStaff(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (s *staffHandler) Login(ctx *gin.Context) {
	req := new(model.StaffLoginRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := s.staffService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
