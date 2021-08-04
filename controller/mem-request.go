package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemRequest(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.MemRequest(), nil)
}
