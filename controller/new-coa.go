package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func NewCOA(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.NewCOA(), nil)
}
