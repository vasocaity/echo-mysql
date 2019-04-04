package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	user_handler "github.com/vasocaity/echo-mysql/user_handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Route => handler
	u := user_handler.NewHandler()
	g := e.Group("/user")
	g.GET("", u.GetValue)

	user_handler.NewHttpHandler(g)

	e.Logger.Fatal(e.Start(":8000"))
}
