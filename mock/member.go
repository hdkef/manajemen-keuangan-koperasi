package mock

import "manajemen-keuangan-koperasi/models"

func Member() models.Member {
	return models.Member{
		User:       models.User{},
		CurBalance: "250000",
		CurDebt:    "400000",
		RecentTransaction: []models.MemTransaction{
			{
				ID:     "1",
				Date:   "4 Aug 21",
				Name:   "Debt Minus",
				Amount: "500000",
			},
			{
				ID:     "2",
				Date:   "4 Aug 21",
				Name:   "Deposit",
				Amount: "500000",
			},
			{
				ID:     "3",
				Date:   "4 Aug 21",
				Name:   "Withdraw",
				Amount: "200000",
			},
			{
				ID:     "4",
				Date:   "1 Jan 21",
				Name:   "Debt Plus",
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
