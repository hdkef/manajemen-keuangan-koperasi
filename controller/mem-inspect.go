package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemInspect(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//if no query params, render mem inspect if there's query params find and render mem inspect res
		_, valid := c.GetQuery(konstanta.QueryID)
		if !valid {
			services.RenderPages(c, HTMLFILENAME.MemInspect(), nil)
			return
		}
		//TOBE get every info available on model Mem-Inspect
		memInfo := mock.MemInspectRes()
		services.RenderPages(c, HTMLFILENAME.MemInspectRes(), memInfo)
	}
}
