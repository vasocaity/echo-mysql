package models

import "time"

type Area struct {
	ID            int32     `json:"id"`
	AreaName      string    `json:"area_name"`
	AreaManagerID int32     `json:"area_manager_id"`
	Code          string    `json:"code"`
	Created       time.Time `json:"created"`
	Modified      time.Time `json:"modified"`
	CreateUID     int32     `json:"create_uid"`
	UpdateUID     int32     `json:"update_uid"`
}
type Areas struct {
	Areas []Area `json:"areas"`
}
