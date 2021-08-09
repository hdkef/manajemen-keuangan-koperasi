package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func NewCOA(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.NewCOA(), nil)
}
