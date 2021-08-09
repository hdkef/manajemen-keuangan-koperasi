package controller

import (
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Member(c *gin.Context) {
	//TOBE get all info on member
	member := mock.Member()

	services.RenderPages(c, HTMLFILENAME.Member(), member)
}
