package api

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v *echo.Group) {
	v.GET("/hello", h.Hello)
	plantGroup := v.Group("/plants")
	plantGroup.POST("/add", h.AddPlant)
	plantGroup.POST("/update", h.UpdatePlant)
	plantGroup.DELETE("/delete", h.DeletePlant)
}
