package main

import (
	"github.com/Aoi-Avant/go-nuxt-tutorial/databases"
	"github.com/Aoi-Avant/go-nuxt-tutorial/models"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
