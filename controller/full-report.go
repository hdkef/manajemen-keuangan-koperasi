package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func FullReport(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.FullReport(), nil)
}
