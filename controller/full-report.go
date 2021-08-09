package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func FullReport(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.FullReport(), nil)
}
