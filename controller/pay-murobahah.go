package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func PayMurobahah(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		services.RenderPages(c, HTMLFILENAME.PayMurobahah(), nil)
	}
}
