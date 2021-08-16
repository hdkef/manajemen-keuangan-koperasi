package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MurobahahReq(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		services.RenderPages(c, HTMLFILENAME.MurobahahReq(), nil)
	}
}
