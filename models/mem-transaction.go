package models

import "database/sql"

type MemTransaction struct {
	ID         float64
	Date       string
	Type       string
	Amount     string
	Info       sql.NullString
	ApprovedBy float64
}
