package mock

import "manajemen-keuangan-koperasi/models"

func Summary() []models.Ledger {
	return []models.Ledger{
		{
			Code:   "A1",
			Name:   "Cash",
			Type:   "Assets",
			Debit:  "456000000",
			Credit: "0",
			Diff:   "456000000",
		},
		{
			Code:   "A2",
			Name:   "Debt",
			Type:   "Liabilities",
			Debit:  "156000000",
			Credit: "0",
			Diff:   "156000000",
		},
	}
}
