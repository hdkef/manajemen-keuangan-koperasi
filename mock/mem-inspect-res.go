package mock

import (
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
)

func MemInspectRes() models.MemInspect {
	return models.MemInspect{
		User: models.User{
			ID:       1,
			Username: "hdk34",
			FullName: "Hadekha Erfadila Fitra",
			DOB:      "20 Jan 1987",
			Tel:      "0812131415116",
			Address:  "Jl. Cikudapateuh",
			Role:     konstanta.RoleADMINInput,
		},
		Balance:              "2500000",
		TotalLoan:            "1500000",
		RemainingLoan:        "500000",
		TotalTransaction:     "23",
		TransactionThisMonth: "2",
		Transaction: []models.MemTransaction{
			{
				Date:   "4 Aug 21",
				Amount: "500000",
			},
			{
				Date:   "4 Aug 21",
				Amount: "500000",
			},
			{
				Date:   "4 Aug 21",
				Amount: "200000",
			},
			{
				Date:   "1 Jan 21",
				Amount: "1000000",
			},
		},
		DebtDetail: []models.DebtDetail{
			{
				ID:        "213",
				Date:      "1 Jan 21",
				DueDate:   "1 Dec 21",
				Amount:    "2500000",
				Remaining: "400000",
				Info:      "beli motor",
			},
		},
	}
}
