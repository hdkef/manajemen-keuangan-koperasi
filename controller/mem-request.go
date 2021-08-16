package controller

import (
	"context"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func MemRequest(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {

		//Begin Transaction
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//get all member request
		memreqs, err := DB.FindMemReqTx(tx)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//get all member murobahah request

		murobahahs, err := DB.FindMemReqMurobahahTx(tx)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		var data map[string]interface{} = make(map[string]interface{})
		data["MemReq"] = memreqs
		data["MemReqMurobahah"] = murobahahs

		services.RenderPages(c, HTMLFILENAME.MemRequest(), data)
	}
}
