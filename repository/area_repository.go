package repository

import (
	"database/sql"
	"fmt"

	models "github.com/vasocaity/echo-mysql/models"
)

var con *sql.DB

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAreas() *models.Areas {
	// con := db.ConnectDB()
	// Execute the query
	var area models.Area
	err := r.db.QueryRow("SELECT * FROM areas").Scan(&area.ID, &area.AreaName, &area.AreaManagerID, &area.Code, &area.Created, &area.Modified, &area.CreateUID, &area.UpdateUID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	result := models.Areas{}
	area = models.Area{}
	result.Areas = append(result.Areas, area)
	return &result
}

func (r *repository) GetArea() (*models.Areas, error) {
	// con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT * FROM areas"

	rows, err := r.db.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	result := new(models.Areas)

	for rows.Next() {
		area := models.Area{}

		err2 := rows.Scan(&area.ID, &area.AreaName, &area.AreaManagerID, &area.Code, &area.Created, &area.Modified, &area.CreateUID, &area.UpdateUID)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		// fmt.Println(area.Created.Format(time.RFC3339) + "\n")
		// area.Created = area.Created.Format(time.RFC3339)
		result.Areas = append(result.Areas, area)
	}
	return result, nil

}
func (r *repository) GetAreaByID(id string) (*models.Area, error) {
	// con := db.CreateCon()
	area := new(models.Area)
	err := r.db.QueryRow("SELECT * FROM areas WHERE id = ?", id).Scan(&area.ID, &area.AreaName, &area.AreaManagerID, &area.Code, &area.Created, &area.Modified, &area.CreateUID, &area.UpdateUID)
	if err != nil {
		return nil, err
	}
	return area, nil
}
