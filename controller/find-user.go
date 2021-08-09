package controller

import (
	"fmt"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func FindUser(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		filter, _ := c.GetQuery(konstanta.QueryFilter)
		key, _ := c.GetQuery(konstanta.QueryKey)
		fmt.Println(filter, key)
		usr := mock.FindUser()
		services.RenderPages(c, HTMLFILENAME.FindUser(), usr)
	}
}
