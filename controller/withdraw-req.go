package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func WithdrawReq(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.WithdrawReq(), nil)
}
