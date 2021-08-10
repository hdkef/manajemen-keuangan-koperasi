package models

type MemReq struct {
	ID      interface{}
	User    User
	Date    string
	Type    string
	Amount  float64
	Doc     string
	DueDate string
	Info    string
}
