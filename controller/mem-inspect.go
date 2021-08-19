package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MemInspect(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {

		//if no UID, render with nil data
		uidStr, valid := c.GetQuery(konstanta.QueryUID)
		if !valid {
			services.RenderPages(c, HTMLFILENAME.MemInspect(), nil)
			return
		}

		//if there is UID, do the same like member controller
		uid, err := strconv.ParseFloat(uidStr, 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		member, err := getMemberFromDB(DB, c, uid)
		if err != nil {
			return
		}
		services.RenderPages(c, HTMLFILENAME.MemInspect(), member)
	}
}
