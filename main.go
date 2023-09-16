package main

import (
	"server-golang/configs"
	helper "server-golang/helpers"
	"server-golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load Env
	helper.LoadEnv()

	// Inisialisasi Database
	configs.Init()

	// Instance framework echo
	app := echo.New()

	// Routing
	routes.Init(app)

	// Hide Banner Echo
	app.HideBanner = true

	// Start server
	app.Start(helper.GetPort())

}
