package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemInspect(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.MemInspect(), nil)
}
