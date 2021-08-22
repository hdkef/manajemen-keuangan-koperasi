package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(C *driver.RedisDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get member uid
		uid, _ := c.Params.Get(konstanta.QueryUID)
		if err := doLogout(uid, C, c); err != nil {
			utils.ResErr(c, http.StatusInternalServerError, &err)
			return
		}
		utils.ResJSON(c, "OK")
	}
}

func doLogout(uid string, C *driver.RedisDriver, c *gin.Context) error {
	//delete member cache in redis
	err := C.Del(uid)
	if err != nil {
		return err
	}
	//delete cookies
	c.SetCookie(konstanta.CookiesBearer, "", 0, "/", "", false, false)
	return nil
}
