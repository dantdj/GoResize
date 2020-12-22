package routes

import (
	"github.com/dantdj/GoResize/internal/handlers"
	"github.com/labstack/echo/v4"
)

func UseRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/health", handlers.HealthCheckHandler)
	e.POST("/resize", handlers.ResizeHandler)

	return e
}
