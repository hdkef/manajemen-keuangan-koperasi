package middleware

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"

	"github.com/gin-gonic/gin"
)

func DB(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set(konstanta.DB, DB)
		c.Next()
	}
}
