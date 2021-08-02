package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func EditCOA(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.EditCOA(), nil)
}
