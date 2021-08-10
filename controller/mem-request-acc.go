package controller

import (
	"context"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MemRequestAcc(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		type_ := c.PostForm(konstanta.QueryType)
		amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		info := c.PostForm(konstanta.QueryInfo)
		claims, _ := c.Get(konstanta.Claims)
		id, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		//start transaction
		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			RenderError(c, err)
			return
		}

		//first insert to member_journal
		_, err = DB.InsertMemJournalTx(tx, uid, type_, amount, info, claims.(models.User).ID)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//then modify member balance with switch case
		_, err = DB.ModifyMemBalanceTx(tx, type_, amount, uid)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//then delete corresponding mem req
		_, err = DB.DeleteMemReqTx(tx, id)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//commit transaction
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		RenderSuccess(c, "transaction succeded")

	}
}
