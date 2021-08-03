package controller

import (
	"errors"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	if c.MustGet(konstanta.Claims).(models.User).Role != konstanta.RoleADMIN {
		RenderError(c, errors.New("USER IS NOT ADMIN"))
		return
	}
	services.RenderPages(c, HTMLFILENAME.Admin(), nil)
}
