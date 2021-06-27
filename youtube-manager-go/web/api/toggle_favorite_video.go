package api

import (
	"errors"

	"firebase.google.com/go/auth"
	"github.com/Aoi-Avant/go-nuxt-tutorial/middlewares"
	"github.com/Aoi-Avant/go-nuxt-tutorial/models"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type ToggleFavoriteVideoResponse struct {
	VideoId    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoId := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}
		err := dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).
			First(&user).
			Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorite := false
		err = dbs.DB.Table("favorites").
			Where(models.Favorite{UserID: user.ID, VideoID: videoId}).
			First(&favorite).
			Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			favorite = models.Favorite{UserID: user.ID, VideoID: videoId}
			dbs.DB.Create(&favorite)
			isFavorite = true
		} else {
			dbs.DB.Delete(&favorite)
		}

		res := ToggleFavoriteVideoResponse{
			VideoId:    videoId,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
