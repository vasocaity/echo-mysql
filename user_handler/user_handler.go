package user_handler

import (
	"net/http"

	"github.com/labstack/echo"
	r "github.com/vasocaity/echo-mysql/repository"
	s "github.com/vasocaity/echo-mysql/services"
)

type handler struct {
	service s.Service
}

func NewHandler(repo r.Repository) Handler {
	service := s.NewService(repo)
	return &handler{service}
}

func NewHttpHandler(g *echo.Group) {
	g.GET("/vara", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hi")
	})
}

func (h *handler) GetAreas(c echo.Context) error {
	list, err := h.service.GetArea()
	if err != nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, list)
}
func (h *handler) GetAreaByID(c echo.Context) error {
	id := c.Param("id")
	list, err := h.service.GetAreaByID(id)
	if err != nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, list)
}
