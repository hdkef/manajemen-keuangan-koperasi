package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func EditCOA(DB *driver.DBDriver) func(c *gin.Context) {

	return func(c *gin.Context) {
		//TOBE get COA list
		coa := mock.EditCoa()

		services.RenderPages(c, HTMLFILENAME.EditCOA(), coa)
	}
}
