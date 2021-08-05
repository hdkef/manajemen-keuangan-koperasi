package controller

import (
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemRequest(c *gin.Context) {
	//TOBE get all member request
	memreq := mock.MemReq()
	services.RenderPages(c, HTMLFILENAME.MemRequest(), memreq)
}
