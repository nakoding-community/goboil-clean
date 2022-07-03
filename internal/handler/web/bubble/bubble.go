package bubble

import (
	"net/http"

	"github.com/nakoding-community/goboil-clean/internal/factory"

	"github.com/labstack/echo/v4"
)

type handler struct {
	Factory factory.Factory
}

func NewHandler(f factory.Factory) *handler {
	return &handler{f}
}

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
}

func (h *handler) Get(c echo.Context) error {
	type M map[string]interface{}
	data := M{"message": "Hello World!"}
	return c.Render(http.StatusOK, "bubble.html", data)
}
