package main

import (
	"flight_system/internal/global"
	"flight_system/internal/interfaces/api/router"
	"flight_system/internal/interfaces/api/server"
)

func main() {
	// init logger
	global.InitLog()

	// init database
	global.InitDatabase()

	// init redis
	global.InitRedis()

	// init server
	server := initServer()

	// run server
	server.Run()
}

func initServer() server.Server {
	// init router
	router := router.NewRouter()

	// init server
	server := server.NewServer("8080", global.Log)
	server.SetShutdownHandler(func() {
		global.Log.Info("[Runtime] Server is shutting down")
		global.Shutdown()
	})
	server.RegisterDefaultCORS()
	server.RegisterRouter(router)

	return server
}
