package routes

import (
	"github.com/Aoi-Avant/go-nuxt-tutorial/web/api"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}
}
