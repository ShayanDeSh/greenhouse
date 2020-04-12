package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello\n")
}
