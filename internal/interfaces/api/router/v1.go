package router

import "flight_system/internal/interfaces/api/middleware"

func (r *router) RegisterV1(group Versions) {
	for _, version := range group.versions {
		// function layer
		user := version.Group("/user")
		airplane := version.Group("/airplane")
		airport := version.Group("/airport")
		flight := version.Group("/flight")
		ticket := version.Group("/ticket")
		order := version.Group("/order")

		// user layer
		{
			user.POST("/sign_up", middleware.UserAuthMiddlerware(), r.userHandler.SignUp)
			user.POST("/login", middleware.UserAuthMiddlerware(), r.userHandler.Login)
		}

		// airplane layer
		{
			airplane.POST("/create", middleware.StaffAuthMiddlerware(), r.airplaneHandler.Create)
			airplane.GET("/get", middleware.StaffAuthMiddlerware(), r.airplaneHandler.Get)
		}

		// airport layer
		{
			airport.POST("/create", middleware.StaffAuthMiddlerware(), r.airportHandler.Create)
			airport.GET("/get", middleware.StaffAuthMiddlerware(), r.airportHandler.Get)
		}

		// flight layer
		{
			flight.POST("/create", middleware.StaffAuthMiddlerware(), r.flightHandler.Create)
			flight.GET("/get", middleware.StaffAuthMiddlerware(), r.flightHandler.GetById)
			flight.GET("/search", r.flightHandler.Search)
		}

		// ticket layer
		{
			ticket.POST("/create_order", middleware.UserAuthMiddlerware(), r.ticketHandler.CreateTicketOrder)
			ticket.GET("/get", middleware.UserAuthMiddlerware(), r.ticketHandler.GetById)
			ticket.GET("/get_by_user/:id", middleware.UserAuthMiddlerware(), r.ticketHandler.GetByUserId)
		}

		// order layer
		{
			order.POST("/create", middleware.UserAuthMiddlerware(), r.orderHandler.CreateOrder)
		}
	}
}
