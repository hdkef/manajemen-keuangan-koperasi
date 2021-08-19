package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"

	"github.com/gin-gonic/gin"
)

func Member(DB *driver.DBDriver, C *driver.RedisDriver) func(c *gin.Context) {
	return func(c *gin.Context) {

		claims, exist := c.Get(konstanta.Claims)
		if !exist {
			RenderError(c, errors.New("auth failed"))
			return
		}
		uid := claims.(models.User).ID

		//try get member from redis and render it
		err := tryRenderWithRedis(c, C, uid)
		if err == nil {
			return
		}

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

		//get murobahahs
		murobahahs, err := DB.FindMemMurobahahTx(tx, uid)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//get all info

		allinfo, err := DB.FindAllInfoTx(tx, uid)
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

		member := models.Member{
			User:              user,
			Balance:           balance,
			RecentTransaction: journal,
			AllInfo:           allinfo,
			Murobahah:         murobahahs,
		}

		services.RenderPages(c, HTMLFILENAME.Member(), member)

		//after get from database success, store it to redis
		err = storeToRedis(c, C, uid, member)
		if err != nil {
			log.Println(err)
		}
	}
}

func tryRenderWithRedis(c *gin.Context, C *driver.RedisDriver, uid float64) error {
	uidInString := fmt.Sprint(uid)

	member, err := C.GetCacheMember(uidInString)
	if err != nil {
		return err
	}

	services.RenderPages(c, HTMLFILENAME.Member(), member)
	return nil
}

func storeToRedis(c *gin.Context, C *driver.RedisDriver, uid float64, member models.Member) error {
	uidInString := fmt.Sprint(uid)

	return C.SetCacheMember(uidInString, member)
}
