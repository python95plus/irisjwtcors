package main

import (
	"irisweb25/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func NewApp() *iris.Application {
	models.Register()
	models.Db.AutoMigrate(
		&models.User{},
		&models.OauthToken{},
	)

	iris.RegisterOnInterrupt(func() {
		_ = models.Db
	})
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}

func main() {
	app := NewApp()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
