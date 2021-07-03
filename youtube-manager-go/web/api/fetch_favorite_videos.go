package api

import (
	"firebase.google.com/go/auth"
	"github.com/Aoi-Avant/go-nuxt-tutorial/middlewares"
	"github.com/Aoi-Avant/go-nuxt-tutorial/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchFavoriteVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		user := models.User{}
		dbs.DB.Table("users").Where(models.User{UID: token.UID}).First(&user)

		favorites := []models.Favorite{}
		dbs.DB.Model(&user).Association("Favorites").Find(&favorites)

		videoIds := ""
		for _, f := range favorites {
			if len(videoIds) == 0 {
				videoIds += f.VideoID
			} else {
				videoIds += "," + f.VideoID
			}
		}

		yts := c.Get("yts").(*youtube.Service)
		call := yts.Videos.
			List([]string{"id", "snippet"}).
			Id(videoIds).
			MaxResults(10)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
