package models

type Member struct {
	User              User
	Balance           MemBalance
	RecentTransaction []MemTransaction
	AllInfo           []AllInfo
	Murobahah         []MemMurobahah
}
