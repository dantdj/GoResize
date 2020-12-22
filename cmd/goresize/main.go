package main

import (
	"github.com/dantdj/GoResize/internal/routes"
)

func main() {
	serverRoutes := routes.ServerRoutes()
	serverRoutes.Logger.Fatal(serverRoutes.Start(":8080"))
}
