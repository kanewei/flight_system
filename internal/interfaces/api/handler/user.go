package handler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userHandler struct {
	userService *service.UserService
}

func NewUserHandler() UserHandler {
	return &userHandler{
		userService: service.NewUserService(),
	}
}

func (u *userHandler) SignUp(ctx *gin.Context) {
	req := new(model.SignUpRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	// Todo validate email

	id, err := u.userService.SignUp(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (u *userHandler) Login(ctx *gin.Context) {
	req := new(model.LoginRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	id, err := u.userService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
