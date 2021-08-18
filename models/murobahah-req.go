package models

type MurobahahReq struct {
	ID      float64
	Agent   User
	Buyer   User
	Amount  float64
	Date    string
	DueDate string
	Doc     string
	Info    string
}
