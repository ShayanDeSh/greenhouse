package api

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8000")

	e.GET("/v1/hello").
		Expect().
		Status(http.StatusOK)

	plant := map[string]interface{}{
		"Ip":         "192.168.2.125",
		"Humidity":   14,
		"NeedWater":  true,
		"Threshhold": 10,
	}
	obj := e.POST("/v1/plants/add").
		WithJSON(plant).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object()
	id := obj.Value("ID").Number().Raw()

	plant["NeedwWater"] = false
	plant["ID"] = id
	e.POST("/v1/plants/update").
		WithJSON(plant).
		Expect().
		Status(http.StatusOK)

	e.DELETE("/v1/plants/delete").
		WithJSON(map[string]interface{}{"id": id}).
		Expect().
		Status(http.StatusOK)
}
