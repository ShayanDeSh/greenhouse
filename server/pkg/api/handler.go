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

func (h *Handler) AddPlant(c echo.Context) (err error) {
	p := new(models.Plant)
	if err = c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err = h.plants.Add(p); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) UpdatePlant(c echo.Context) (err error) {
	p := new(models.Plant)

	if err = c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err = h.plants.Update(p); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handler) DeletePlant(c echo.Context) (err error) {
	m := echo.Map{}
	if err = c.Bind(&m); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	id := int(m["id"].(float64))
	if err = h.plants.Delete(id); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Deleted\n")
}
