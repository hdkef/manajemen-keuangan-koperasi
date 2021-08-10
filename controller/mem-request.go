package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemRequest(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TOBE get all member request

		memreqs, err := DB.FindMemReq()
		if err != nil {
			RenderError(c, err)
			return
		}

		services.RenderPages(c, HTMLFILENAME.MemRequest(), memreqs)
	}
}
