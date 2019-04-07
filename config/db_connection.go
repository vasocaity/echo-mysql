package config

import (
	"database/sql"
	"fmt"
)

type DbConn struct {
	db *sql.DB
}

func ConnectDB() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/test-plantd")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connect DB success")
	// perform a db.Query insert
	// insert, err := db.Query("INSERT INTO areas VALUES (1,'vara',1,'001','2012-06-18T10:34:09','2012-06-18T10:34:09',1,1)")

	// if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// be careful deferring Queries if you are using transactions
	// defer insert.Close()

	// var area models.Area
	// Execute the query
	// err = db.QueryRow("SELECT * FROM areas").Scan(&area.ID, &area.AreaName, &area.AreaManagerID, &area.Code, &area.Created, &area.Modified, &area.CreateUID, &area.UpdateUID)
	// if err != nil {
	// panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// fmt.Println(area.ID)
	// fmt.Println(area.AreaName)
	// fmt.Println(area.Code)
}

/*Create mysql connection*/
func CreateCon() (*sql.DB, error) {

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
	return db, err
}
