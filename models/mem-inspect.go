package models

type MemInspect struct {
	User                 User
	Balance              string
	TotalLoan            string
	RemainingLoan        string
	TotalTransaction     string
	TransactionThisMonth string
	Transaction          []MemTransaction
	DebtDetail           []DebtDetail
}
