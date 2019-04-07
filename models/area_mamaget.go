package models

import "time"

type AreaManager struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	LineID    string    `json:"line_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
	CreateUID int32     `json:"create_uid"`
	UpdateUID int32     `json:"update_uid"`
}
