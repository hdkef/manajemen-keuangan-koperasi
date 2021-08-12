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
	"time"

	"github.com/gin-gonic/gin"
)

func WithdrawReq(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			claims, exist := c.Get(konstanta.Claims)
			if !exist {
				RenderError(c, errors.New("claims not exist"))
				return
			}

			uid := claims.(models.User).ID

			amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
			if err != nil {
				RenderError(c, err)
				return
			}
			info := c.PostForm(konstanta.QueryInfo)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			//begin transaction
			tx, err := DB.DB.BeginTx(ctx, nil)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//find Sumbangan Sukarela balance
			ss, err := DB.FindMemSSBalanceTx(tx, uid)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			// ss must greater or equal with amount of withdraw
			if ss < amount {
				tx.Rollback()
				RenderError(c, errors.New("not enough balance"))
				return
			}

			//send request

			_, err = DB.InsertMemReqTx(tx, uid, konstanta.TypeSSNeg, amount, driver.MemReqOption{
				DueDate: time.Now(),
				Info:    info,
			})
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//render success page
			RenderSuccess(c, "withdraw request created")
			return
		}
		services.RenderPages(c, HTMLFILENAME.WithdrawReq(), nil)
	}
}
