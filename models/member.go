package models

type Member struct {
	User              User
	CurBalance        string
	CurDebt           string
	RecentTransaction []MemTransaction
	DebtDetail        []DebtDetail
}
