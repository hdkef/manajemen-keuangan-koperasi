package controller

import (
	"manajemen-keuangan-koperasi/common"
	"manajemen-keuangan-koperasi/konstanta"

	"github.com/gin-gonic/gin"
)

var HTMLFILENAME = konstanta.GetHTMLFileName()

func Home(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.Home(), nil)
}
