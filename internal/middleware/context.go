package middleware

import (
	"context"

	"github.com/nakoding-community/goboil-clean/internal/abstraction"
	"github.com/nakoding-community/goboil-clean/pkg/constant"

	"github.com/labstack/echo/v4"
)

func Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &abstraction.Context{
			Context: c.Request().Context(),
		}

		newRequest := c.Request().WithContext(context.WithValue(c.Request().Context(), constant.CONTEXT_KEY, cc))
		c.SetRequest(newRequest)

		return next(c)
	}
}
