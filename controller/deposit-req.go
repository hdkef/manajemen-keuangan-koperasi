package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func DepositReq(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.DepositReq(), nil)
}
