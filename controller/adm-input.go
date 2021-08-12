package controller

import (
	"context"
	"errors"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdmInput(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {

			uid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
			if err != nil {
				RenderError(c, err)
				return
			}
			amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
			if err != nil {
				RenderError(c, err)
				return
			}
			type_ := c.PostForm(konstanta.QueryType)
			info := c.PostForm(konstanta.QueryInfo)
			claims, exist := c.Get(konstanta.Claims)
			if !exist {
				RenderError(c, errors.New("auth failed"))
				return
			}
			approvedby := claims.(models.User).ID

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			//begin transaction
			tx, err := DB.DB.BeginTx(ctx, nil)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//modify balance
			_, err = DB.ModifyMemBalanceTx(tx, type_, amount, uid)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//add to journal
			_, err = DB.InsertMemJournalTx(tx, uid, type_, amount, info, approvedby)
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

			RenderSuccess(c, "transaction succeeded")
			return
		}

		services.RenderPages(c, HTMLFILENAME.AdmInput(), nil)
	}
}
