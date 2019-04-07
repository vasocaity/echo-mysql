package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	db "github.com/vasocaity/echo-mysql/config"
	repo "github.com/vasocaity/echo-mysql/repository"
	user_handler "github.com/vasocaity/echo-mysql/user_handler"
)

// var db *sql.DB

func main() {
	// connect db
	db, err := db.CreateCon()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// create handler
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	repo := repo.NewRepository(db)
	u := user_handler.NewHandler(repo)
	g := e.Group("/user")
	g.GET("", u.GetValue)
	g.GET("/:id", u.GetAreaByID)
	user_handler.NewHttpHandler(g)

	e.Logger.Fatal(e.Start(":8000"))
}
