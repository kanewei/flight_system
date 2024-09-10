package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler interface {
	Check(ctx *gin.Context)
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return new(healthHandler)
}

func (h *healthHandler) Check(c *gin.Context) {
	c.Status(http.StatusOK)
}
