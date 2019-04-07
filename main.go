package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	repo "github.com/vasocaity/echo-mysql/repository"
	user_handler "github.com/vasocaity/echo-mysql/user_handler"
)

var db *sql.DB

func main() {
	// connect db
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/test-plantd?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
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
