package api

import (
	"errors"

	"firebase.google.com/go/auth"
	"github.com/Aoi-Avant/go-nuxt-tutorial/middlewares"
	"github.com/Aoi-Avant/go-nuxt-tutorial/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
)

type VideoResponse struct {
	VideoList  *youtube.VideoListResponse `json:"video_list"`
	IsFavorite bool                       `json:"is_favorite"`
}

func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		videoId := c.Param("id")

		isFavorite := false
		if token != nil {
			favorite := models.Favorite{}
			isFavoriteNotFound := dbs.DB.Table("favorites").
				Joins("INNER JOIN users ON users.id = favorites.user_id").
				Where(models.User{UID: token.UID}).
				Where(models.Favorite{VideoID: videoId}).
				First(&favorite).
				Error

			logrus.Debug("isFavoriteNotFound: ", isFavoriteNotFound)

			if errors.Is(isFavoriteNotFound, gorm.ErrRecordNotFound) {
				isFavorite = true
			}
		}

		call := yts.Videos.
			List([]string{"id", "snippet"}).
			Id(videoId)

		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error calling YouTube API %v", err)
		}

		v := VideoResponse{
			VideoList:  res,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, v)
	}
}
