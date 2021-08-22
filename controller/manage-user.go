package controller

import (
	"context"
	"fmt"
	"log"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ManageUser(DB *driver.DBDriver, C *driver.RedisDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			uid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUID), 64)
			if err != nil {
				RenderError(c, err)
				return
			}
			isagent := c.PostForm(konstanta.QueryIsAgent)
			role := c.PostForm(konstanta.QueryRole)

			//Begin transaction
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			tx, err := DB.DB.BeginTx(ctx, nil)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//if isagent is not default
			if isagent != "Default" {
				_, err = DB.ModifyMemberFieldByUIDTx(tx, uid, konstanta.TYPEIsAgent, isagent)
				if err != nil {
					tx.Rollback()
					RenderError(c, err)
					return
				}
			}

			//if role is not default
			if role != "Default" {
				_, err = DB.ModifyMemberFieldByUIDTx(tx, uid, konstanta.TYPERole, role)
				if err != nil {
					tx.Rollback()
					RenderError(c, err)
					return
				}
			}

			//dont forget to commit

			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//set redis cache for force logout
			err = C.SetString(konstanta.ForceLogoutKey(uid), "Y")
			if err != nil {
				log.Println(err)
			}

			//Render Sucess
			RenderSuccess(c, fmt.Sprintf("success modify role %s isagent %s", role, isagent))
			return
		}
		services.RenderPages(c, HTMLFILENAME.ManageUser(), nil)
	}
}
