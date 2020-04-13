package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nothingrealm/greenhouse/server/pkg/db/models"
	"net/http"
)

type Handler struct {
	plants *models.Plants
}

func NewHandler(plants *models.Plants) *Handler {
	return &Handler{plants}
}

func (h *Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello\n")
}
