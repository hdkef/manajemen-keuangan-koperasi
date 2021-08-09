package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Summary(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TOBE get all info on summary
		sum := mock.Summary()
		services.RenderPages(c, HTMLFILENAME.Summary(), sum)
	}
}
