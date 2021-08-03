package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func renderError(c *gin.Context, err error) {
	services.RenderPages(c, HTMLFILENAME.Error(), err.Error())
}
