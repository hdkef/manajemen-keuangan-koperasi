package models

type MemReq struct {
	ID     float64
	User   User
	Date   string
	Type   string
	Amount float64
	Info   string
}
