package controller

import (
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

var HTMLFILENAME = konstanta.GetHTMLFileName()
var route = konstanta.GetRoute()

func Home(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.Home(), nil)
}
