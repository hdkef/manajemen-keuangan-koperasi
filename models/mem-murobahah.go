package models

import "database/sql"

type MemMurobahah struct {
	ID      float64
	Initial float64
	Paid    float64
	Date    string
	DueDate string
	Doc     string
	Info    sql.NullString
}
