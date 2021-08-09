package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func FullReport(DB *driver.DBDriver) func(*gin.Context) {
	return func(c *gin.Context) {
		services.RenderPages(c, HTMLFILENAME.FullReport(), nil)
	}
}
