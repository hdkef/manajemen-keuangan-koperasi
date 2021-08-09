package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func UserSetting(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.UserSetting(), nil)
}
