package controller

import (
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func NewTransaction(c *gin.Context) {
	services.RenderPages(c, HTMLFILENAME.NewTransaction(), nil)
}
