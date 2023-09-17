package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	auth "server-golang/controllers/auth"
	mid "server-golang/middleware"
)

func Init(app *echo.Echo) {
	app.Use(middleware.Logger())
	// Tanpa validasi auth
	app.POST("/login", auth.Login)
	app.POST("/register", auth.Register)
	app.POST("/logout", auth.Logout)

	// Endpoint perlu validasi auth
	app.POST("/api/users/profile", mid.JwtAuth(auth.AddProfile))
	app.PUT("/api/users/profile", mid.JwtAuth(auth.UpdateProfile))

}
