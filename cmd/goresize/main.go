package main

import (
	"github.com/dantdj/GoResize/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e = routes.UseRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
