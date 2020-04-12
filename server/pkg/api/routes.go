package api

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v *echo.Group) {
	v.GET("/hello", h.Hello)
}