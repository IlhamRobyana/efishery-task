package route

import (
	"github.com/ilhamrobyana/efishery-task/auth"
	"github.com/labstack/echo"
)

func authRoute(e *echo.Echo) {
	g := e.Group("/auth")
	g.POST("/register", auth.Register)
	g.POST("/login", auth.Login)
	g.GET("/validate-token", auth.ValidateToken)
}
