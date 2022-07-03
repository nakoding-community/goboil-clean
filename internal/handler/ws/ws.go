package ws

import (
	"github.com/nakoding-community/goboil-clean/internal/factory"
	"github.com/nakoding-community/goboil-clean/internal/handler/ws/course"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "ws"
	course.NewHandler(f).Route(e.Group(prefix + "/course"))
}
