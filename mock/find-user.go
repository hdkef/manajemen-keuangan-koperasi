package mock

import (
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
)

func FindUser() []models.User {
	return []models.User{
		{
			ID:       1,
			Username: "hdkef",
			FullName: "Hadekha Erfadila Fitra",
			Role:     konstanta.RoleADMINInput,
		},
		{
			ID:       1,
			Username: "dsz",
			FullName: "Diaz Zulfikar",
			Role:     konstanta.RoleMEMBER,
		},
	}
}
