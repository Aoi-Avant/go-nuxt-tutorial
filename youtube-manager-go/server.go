package main

import (
	"github.com/Aoi-Avant/go-nuxt-tutorial/middlewares"
	"github.com/Aoi-Avant/go-nuxt-tutorial/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YoutubeService())

	// Routes
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
