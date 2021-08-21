package models

import "database/sql"

type Transaction struct {
	ID     string
	Date   string
	Name   string
	Debit  string
	Credit string
	Info   sql.NullString
}
