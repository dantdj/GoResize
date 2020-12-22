package main

import (
	"github.com/dantdj/GoResize/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e = routes.UseRoutes(e)

	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
