package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, err error) {
	services.RenderPages(c, HTMLFILENAME.Error(), err.Error())
}
