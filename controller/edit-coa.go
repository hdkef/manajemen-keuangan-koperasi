package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func EditCOA(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.EditCOA(), nil)
}
