package main

import (
	"github.com/labstack/echo/v4"
	DB "server-golang/helpers"
	helper "server-golang/helpers"
	"server-golang/routes"
)

func main() {
	// Load Env
	helper.LoadEnv()

	// Inisialisasi Database
	DB.Init()

	// Instance framework echo
	app := echo.New()

	// Routing
	routes.Init(app)

	// Start server
	app.Start(helper.GetPort())

}
