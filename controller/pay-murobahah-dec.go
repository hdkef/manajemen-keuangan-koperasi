package controller

import (
	"manajemen-keuangan-koperasi/driver"

	"github.com/gin-gonic/gin"
)

func PayMurobahahDec(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get all value

		//delete pay req

		//inform user transactin been declined

		//respond
		RenderSuccess(c, "transaction declined")
	}
}
