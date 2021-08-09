package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"net/http"

	"github.com/gin-gonic/gin"
)

var HTMLFILENAME = konstanta.GetHTMLFileName()
var route = konstanta.GetRoute()

func Home(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, route.Login())
	}
}
