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

func MurobahahAcc(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get all values
		claims, exist := c.Get(konstanta.Claims)
		if !exist {
			RenderError(c, errors.New("no claims"))
			return
		}
		uid := claims.(models.User).ID
		id, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		agentid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUIDAgent), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		buyerid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUIDBuyer), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		doc := c.PostForm(konstanta.QueryDoc)
		duedate := c.PostForm(konstanta.QueryDueDate)
		info := c.PostForm(konstanta.QueryInfo)

		//begin transaction
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//insert into journal
		_, err = DB.InsertMemJournalTx(tx, buyerid, konstanta.TYPEMurobahahPos, amount, info, uid)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//insert into table member murobahah
		res, err := DB.InsertMemMurobahahTx(tx, buyerid, amount, doc, duedate, info, uid)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//insert into table agent_record
		_, err = DB.InsertAgentHistoryTx(tx, agentid, float64(lastID))
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//delete req from murobahah req
		_, err = DB.DeleteMemReqMurobahahTx(tx, id)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//dont forget to commit
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		RenderSuccess(c, "murobahah been accepted")
	}
}
