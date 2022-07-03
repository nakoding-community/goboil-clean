package rest

import (
	"fmt"
	"net/http"
	"os"

	docs "github.com/nakoding-community/goboil-clean/docs"
	"github.com/nakoding-community/goboil-clean/internal/factory"
	"github.com/nakoding-community/goboil-clean/internal/handler/rest/auth"
	"github.com/nakoding-community/goboil-clean/internal/handler/rest/user"
	"github.com/nakoding-community/goboil-clean/pkg/constant"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, f factory.Factory) {
	var (
		APP     = os.Getenv(constant.APP)
		VERSION = os.Getenv(constant.VERSION)
		HOST    = os.Getenv(constant.HOST)
		SCHEME  = os.Getenv(constant.SCHEME)
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	prefix := "rest"
	auth.NewHandler(f).Route(e.Group(prefix + "/auth"))
	user.NewHandler(f).Route(e.Group(prefix + "/users"))
}
