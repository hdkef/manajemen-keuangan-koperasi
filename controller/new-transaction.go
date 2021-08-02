package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func NewTransaction(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.NewTransaction(), nil)
}
