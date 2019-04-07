package user_handler

import (
	"net/http"

	"github.com/labstack/echo"
	r "github.com/vasocaity/echo-mysql/repository"
)

type handler struct {
	repo r.Repository
}

func NewHandler(repo r.Repository) Handler {
	return &handler{repo}
}

func NewHttpHandler(g *echo.Group) {
	g.GET("/vara", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hi")
	})
}

func (h *handler) GetValue(c echo.Context) error {
	list, err := h.repo.GetArea()
	if err != nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, list)
}
func (h *handler) GetAreaByID(c echo.Context) error {
	id := c.Param("id")
	list, err := h.repo.GetAreaByID(id)
	if err != nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, list)
}
