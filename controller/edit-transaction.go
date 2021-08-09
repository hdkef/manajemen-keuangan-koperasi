package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func EditTransaction(DB *driver.DBDriver) func(c *gin.Context) {

	return func(c *gin.Context) {
		//TOBE get recent transaction
		transaction := mock.EditTransaction()

		services.RenderPages(c, HTMLFILENAME.EditTransaction(), transaction)
	}
}
