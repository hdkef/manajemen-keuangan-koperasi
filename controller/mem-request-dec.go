package controller

import (
	"context"
	"fmt"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MemRequestDec(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//delete corresponding mem req
		id, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		uid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		type_ := c.PostForm(konstanta.QueryType)
		date := c.PostForm(konstanta.QueryDate)
		amount := c.PostForm(konstanta.QueryAmount)

		//mulai transaksi
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//hapus mem request

		_, err = DB.DeleteMemReqTx(tx, id)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}
		//send info "permintaan X tidak diterima"
		_, err = DB.InsertAllInfo(tx, uid, fmt.Sprintf("permintaan %s sebesar %s tanggal %s ditolak", type_, amount, date))
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//jangan lupa commit
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		RenderSuccess(c, "transaction declined")
	}
}
