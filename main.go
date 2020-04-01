package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/quannv132/qor/admin"
	"github.com/quannv132/qor/models"
)

func main() {
	var db *gorm.DB
	var err error
	if db, err = gorm.Open(
		"postgres",
		"user=quan dbname=qor_tutorial password=123456 sslmode=disable",
	); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize database connection")
	}
	defer db.Close()
	// if err = migrate.Start(db); err != nil {
	// 	logrus.WithError(err).Fatal("Couldn't run migration")
	// }
	db.Debug().AutoMigrate(&models.Product{}, &models.User{})

	r := gin.New()
	a := admin.New(db, "", "secret")
	a.Bind(r)
	r.Run("127.0.0.1:8080")
}
