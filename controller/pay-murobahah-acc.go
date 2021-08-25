package controller

import (
	"manajemen-keuangan-koperasi/driver"

	"github.com/gin-gonic/gin"
)

func PayMurobahahAcc(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get all value

		//increment paid murobahah

		//insert into journal

		//delete pay req

		//respond
		RenderSuccess(c, "transaction accepted")
	}
}
