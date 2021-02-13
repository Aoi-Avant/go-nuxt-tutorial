package databases

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	db, err = gorm.Open("mysql",
		os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	)

	if err = nil {
		logrus.Fatal(err)
	}

	return db, err
}
