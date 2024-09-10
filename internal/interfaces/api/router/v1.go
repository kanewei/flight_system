package router

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
			user.POST("/sign_up", r.userHandler.SignUp)
			user.POST("/login", r.userHandler.Login)
		}

		// airplane layer
		{
			airplane.POST("/create", r.airplaneHandler.Create)
			airplane.GET("/get", r.airplaneHandler.Get)
		}

		// airport layer
		{
			airport.POST("/create", r.airportHandler.Create)
			airport.GET("/get", r.airportHandler.Get)
		}

		// flight layer
		{
			flight.POST("/create", r.flightHandler.Create)
			flight.GET("/get", r.flightHandler.GetById)
			flight.GET("/search", r.flightHandler.Search)
		}

		// ticket layer
		{
			ticket.POST("/create_order", r.ticketHandler.CreateTicketOrder)
			ticket.GET("/get", r.ticketHandler.GetById)
			ticket.GET("/get_by_user/:id", r.ticketHandler.GetByUserId)
		}

		// order layer
		{
			order.POST("/create", r.orderHandler.CreateOrder)
		}
	}
}
