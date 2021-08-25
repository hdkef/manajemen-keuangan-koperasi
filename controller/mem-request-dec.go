package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MemRequestDec(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//delete corresponding mem req
		id, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		_, err = DB.DeleteMemReq(id)
		if err != nil {
			RenderError(c, err)
			return
		}
		//TOBE send info "permintaan X tidak diterima"
		RenderSuccess(c, "transaction declined")
	}
}
