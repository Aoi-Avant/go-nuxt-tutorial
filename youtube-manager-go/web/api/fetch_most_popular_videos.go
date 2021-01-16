package api

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(fasthttp.StatusOK, "Most Popular")
	}
}
