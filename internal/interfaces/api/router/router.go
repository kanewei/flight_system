package router

import (
	"flight_system/internal/interfaces/api/handler"
	"flight_system/internal/interfaces/api/middleware"
	"flight_system/internal/interfaces/api/server"

	"github.com/gin-gonic/gin"
)

type router struct {
	healthHandler   handler.HealthHandler
	staffHandler    handler.StaffHandler
	userHandler     handler.UserHandler
	airplaneHandler handler.AirplaneHandler
	airportHandler  handler.AirportHandler
	flightHandler   handler.FlightHandler
	ticketHandler   handler.TicketHandler
	orderHandler    handler.OrderHandler
}

func NewRouter() server.Router {
	return &router{
		healthHandler:   handler.NewHealthHandler(),
		staffHandler:    handler.NewStaffHandler(),
		userHandler:     handler.NewUserHandler(),
		airplaneHandler: handler.NewAirplaneHandler(),
		airportHandler:  handler.NewAirportHandler(),
		flightHandler:   handler.NewFlightHandler(),
		ticketHandler:   handler.NewTicketHandler(),
		orderHandler:    handler.NewOrderHandler(),
	}
}

func (r *router) Register(engine *gin.Engine) {
	// health check
	engine.GET("/healthz", r.healthHandler.Check)

	// auth
	engine.Use(middleware.AuthHandler())

	{
		// endpoint source
		serverGroup := engine.Group("/api")

		// version
		apiV1 := serverGroup.Group("/v1")

		// version group
		versionGroup := NewGroupsByVersion([]*gin.RouterGroup{apiV1})
		r.RegisterV1(versionGroup)
	}
}

type Versions struct {
	versions []*gin.RouterGroup
}

func NewGroupsByVersion(versions []*gin.RouterGroup) Versions {
	return Versions{
		versions,
	}
}
