package controller

import (
	"manajemen-keuangan-koperasi/common"

	"github.com/gin-gonic/gin"
)

func ManageUser(c *gin.Context) {
	common.RenderPages(c, HTMLFILENAME.ManageUser(), nil)
}
