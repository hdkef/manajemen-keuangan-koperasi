package mock

import "manajemen-keuangan-koperasi/models"

func EditTransaction() []models.Transaction {
	return []models.Transaction{
		{
			ID:     "1",
			Date:   "4 Aug 21",
			Name:   "Cash",
			Debit:  "2500000",
			Credit: "",
		},
		{
			ID:     "2",
			Date:   "4 Aug 21",
			Name:   "Cash",
			Debit:  "2500000",
			Credit: "",
		},
	}
}
