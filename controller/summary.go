package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Summary(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.Summary(), nil)
}
