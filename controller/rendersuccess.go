package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func RenderSuccess(c *gin.Context, msg string) {
	services.RenderPages(c, HTMLFILENAME.Success(), msg)
}
