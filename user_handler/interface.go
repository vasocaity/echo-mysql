package user_handler

import (
	"github.com/labstack/echo"
)

type Handler interface {
	GetAreas(c echo.Context) error
	GetAreaByID(c echo.Context) error
}
