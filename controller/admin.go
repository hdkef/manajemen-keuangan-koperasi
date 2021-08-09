package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.Admin(), nil)
}
