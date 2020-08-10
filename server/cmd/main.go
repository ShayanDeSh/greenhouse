package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nothingrealm/greenhouse/server/pkg/api"
	"github.com/nothingrealm/greenhouse/server/pkg/db"
	"github.com/nothingrealm/greenhouse/server/pkg/db/models"
)

func main() {
	db, err := db.Init()
	p := models.NewPlants(db)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// echo initialization
	e := echo.New()
	e.Use(middleware.Logger())

	//

	// Router initialization
	v1 := e.Group("/v1")
	h := api.NewHandler(p)
	h.Register(v1)

	e.Logger.Fatal(e.Start("0.0.0.0:9000"))
}
