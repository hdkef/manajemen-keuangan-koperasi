package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Member(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.Member(), nil)
}
