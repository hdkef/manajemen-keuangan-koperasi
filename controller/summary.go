package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func Summary(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.Summary(), nil)
}
