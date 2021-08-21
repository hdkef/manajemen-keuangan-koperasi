package models

import "database/sql"

type MurobahahReq struct {
	ID      float64
	Agent   User
	Buyer   User
	Amount  float64
	Date    string
	DueDate string
	Doc     string
	Info    sql.NullString
}
