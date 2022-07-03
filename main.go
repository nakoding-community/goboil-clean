package main

import (
	"os"

	db "github.com/nakoding-community/goboil-clean/database"
	"github.com/nakoding-community/goboil-clean/database/migration"
	"github.com/nakoding-community/goboil-clean/database/seeder"
	"github.com/nakoding-community/goboil-clean/internal/factory"
	"github.com/nakoding-community/goboil-clean/internal/http"
	"github.com/nakoding-community/goboil-clean/internal/middleware"
	"github.com/nakoding-community/goboil-clean/pkg/constant"
	"github.com/nakoding-community/goboil-clean/pkg/cron"
	"github.com/nakoding-community/goboil-clean/pkg/sentry"
	"github.com/nakoding-community/goboil-clean/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv(constant.ENV)
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title goboil-clean
// @version 0.0.1
// @description This is a doc for goboil-clean.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv(constant.PORT)

	// dependency
	db.Init()

	// hook
	migration.Init()
	seeder.Init()
	cron.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.Init()
	http.Init(e, f)

	//sentry
	sentry.Init()

	e.Logger.Fatal(e.Start(":" + PORT))
}
