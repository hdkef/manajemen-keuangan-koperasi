package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func FindUser(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		filter, filterExist := c.GetQuery(konstanta.QueryFilter)
		key, keyExist := c.GetQuery(konstanta.QueryKey)
		if keyExist || filterExist {
			//find user
			users, err := DB.FindAllUserByFilter(filter, key)
			if err != nil {
				//if error show nothing
				services.RenderPages(c, HTMLFILENAME.FindUser(), []models.User{})
				return
			}

			services.RenderPages(c, HTMLFILENAME.FindUser(), users)
			return
		}
		services.RenderPages(c, HTMLFILENAME.FindUser(), []models.User{})
	}
}
