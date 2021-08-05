package mock

import "manajemen-keuangan-koperasi/models"

func EditCoa() []models.COA {
	return []models.COA{
		{
			ID:   "1",
			Code: "A1",
			Name: "Cash",
			Info: "Cash is bla bla bla",
		},
		{
			ID:   "1",
			Code: "A2",
			Name: "Loan",
			Info: "Loan is bla bla bla",
		},
	}
}
