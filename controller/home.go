package controller

import (
	"manajemen-keuangan-koperasi/common"
	"manajemen-keuangan-koperasi/konstanta"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	common.RenderPages(c, konstanta.HomeHTML, nil)
}
