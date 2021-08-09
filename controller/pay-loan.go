package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func PayLoan(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		services.RenderPages(c, HTMLFILENAME.PayLoan(), nil)
	}
}
