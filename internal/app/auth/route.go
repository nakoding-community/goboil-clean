package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/nakoding-community/goboil-clean/internal/middleware"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/login", h.Login, middleware.Context)
	g.POST("/register", h.Register, middleware.Context)
}
