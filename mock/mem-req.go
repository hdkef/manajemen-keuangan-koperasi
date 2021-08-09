package mock

import "manajemen-keuangan-koperasi/models"

func MemReq() []models.MemReq {
	return []models.MemReq{
		{
			ID:       "1",
			Date:     "4 Aug 21",
			FullName: "Hadekha Erfadila Fitra",
			Type:     "Deposit Req",
			Amount:   "500000",
		},
		{
			ID:       "1",
			Date:     "3 Aug 21",
			FullName: "Diaz Zulfikar",
			Type:     "Loan Req",
			Amount:   "5000000",
			Document: "l1l1.pdf",
		},
	}
}
