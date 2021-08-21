package models

import "database/sql"

type User struct {
	ID       float64
	MemID    string
	Username string
	Role     string
	Pass     string
	IsAgent  string
	Tel      sql.NullString
}
