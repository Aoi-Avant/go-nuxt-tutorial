package main

import (
	"github.com/Aoi-Avant/go-nuxt-tutorial/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
