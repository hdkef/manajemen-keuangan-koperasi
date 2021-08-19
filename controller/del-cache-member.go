package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DelCacheMember(C *driver.RedisDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, _ := c.Params.Get(konstanta.QueryUID)
		err := C.DelCacheMember(uid)
		if err != nil {
			utils.ResErr(c, http.StatusInternalServerError, &err)
			return
		}
		utils.ResJSON(c, "OK")
	}
}
