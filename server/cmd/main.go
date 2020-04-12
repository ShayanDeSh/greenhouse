package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nothingrealm/greenhouse/server/pkg/api"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	v1 := e.Group("/v1")

	h := api.NewHandler()
	h.Register(v1)

	e.Logger.Fatal(e.Start(":8000"))
}
