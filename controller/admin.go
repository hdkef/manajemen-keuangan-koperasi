package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.Admin(), nil)
}
