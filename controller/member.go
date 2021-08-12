package controller

import (
	"context"
	"errors"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Member(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {

		claims, exist := c.Get(konstanta.Claims)
		if !exist {
			RenderError(c, errors.New("auth failed"))
			return
		}
		uid := claims.(models.User).ID

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		//begin transaction

		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//Get Member Information real time from DB (not Token)

		user, err := DB.FindOneUserByUIDTx(tx, uid)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//Get Member recent transactions journal limit 5

		journal, err := DB.FindLimitedMemJournalByUIDTx(tx, uid, 5)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//Get Member balance info

		balance, err := DB.FindMemBalanceByUIDTx(tx, uid)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//font forget to commit
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		services.RenderPages(c, HTMLFILENAME.Member(), models.Member{
			User:              user,
			Balance:           balance,
			RecentTransaction: journal,
		})
	}
}
