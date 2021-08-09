package controller

import (
	"manajemen-keuangan-koperasi/konstanta"
	"net/http"

	"github.com/gin-gonic/gin"
)

var HTMLFILENAME = konstanta.GetHTMLFileName()
var route = konstanta.GetRoute()

func Home(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, route.Login())
}
