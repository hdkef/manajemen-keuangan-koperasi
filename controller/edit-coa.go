package controller

import (
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func EditCOA(c *gin.Context) {

	//TOBE get COA list
	coa := mock.EditCoa()

	services.RenderPages(c, HTMLFILENAME.EditCOA(), coa)
}
