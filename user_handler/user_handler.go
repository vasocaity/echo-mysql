package user_handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func NewHttpHandler(g *echo.Group) {
	g.GET("/vara", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hi")
	})
}

func (h *handler) GetValue(c echo.Context) error {
	return c.JSON(http.StatusOK, "VARA")
}
