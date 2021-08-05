package controller

import (
	"manajemen-keuangan-koperasi/mock"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func EditTransaction(c *gin.Context) {

	//TOBE get recent transaction
	transaction := mock.EditTransaction()

	services.RenderPages(c, HTMLFILENAME.EditTransaction(), transaction)
}
