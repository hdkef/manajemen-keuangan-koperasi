package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.Register(), nil)
}
