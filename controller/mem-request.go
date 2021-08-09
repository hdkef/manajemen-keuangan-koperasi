package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemRequest(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TOBE get all member request
		memreq := mock.MemReq()
		services.RenderPages(c, HTMLFILENAME.MemRequest(), memreq)
	}
}
