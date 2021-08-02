package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func EditTransaction(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.EditTransaction(), nil)
}
