package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func EditTransaction(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.EditTransaction(), nil)
}
