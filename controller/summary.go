package controller

import (
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Summary(c *gin.Context) {
	//TOBE get all info on summary
	sum := mock.Summary()
	services.RenderPages(c, HTMLFILENAME.Summary(), sum)
}
