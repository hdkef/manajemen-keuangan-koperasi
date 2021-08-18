package models

type Member struct {
	User              User
	Balance           MemBalance
	CurDebt           string
	RecentTransaction []MemTransaction
	DebtDetail        []DebtDetail
	AllInfo           []AllInfo
}
