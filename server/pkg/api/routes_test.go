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
}
