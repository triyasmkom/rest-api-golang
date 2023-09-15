package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	auth "server-golang/controllers/auth"
)

func Init(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.POST("/login", auth.Login)
	app.POST("/register", auth.Register)
	app.POST("/logout", auth.Logout)

}
