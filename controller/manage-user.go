package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func ManageUser(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.ManageUser(), nil)
}
