package controller

import (
	"context"
	"errors"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PayMurobahahAcc(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get all value
		id, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		murobahahid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryMurobahahID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
		if err != nil {
			RenderError(c, err)
			return
		}

		info := c.PostForm(konstanta.QueryInfo)

		claims, exist := c.Get(konstanta.Claims)
		if !exist {
			RenderError(c, errors.New("unauthorized"))
			return
		}
		approvedBy := claims.(models.User).ID

		//start transaction
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//increment paid murobahah

		uid, err := DB.IncrementPaidMurobahahReturnUIDTx(tx, murobahahid, amount)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//insert into journal

		_, err = DB.InsertMemJournalTx(tx, uid, konstanta.TYPEMurobahahNeg, amount, info, approvedBy)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//delete pay req

		_, err = DB.DeleteMemReqMurobahahTx(tx, id)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//commit
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//respond
		RenderSuccess(c, "transaction accepted")
	}
}
