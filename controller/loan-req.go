package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func LoanReq(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.LoanReq(), nil)
}
