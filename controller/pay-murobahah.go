package controller

import (
	"errors"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PayMurobahah(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {

		claims, exist := c.Get(konstanta.Claims)
		if !exist {
			RenderError(c, errors.New("unauthorized"))
			return
		}
		uid := claims.(models.User).ID

		if c.Request.Method == http.MethodPost {
			//get all value

			murobahahID, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
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

			//insert to mem_murobahah_payreq

			_, err = DB.InsertMemMurobahahPayReq(murobahahID, amount, info)
			if err != nil {
				RenderError(c, err)
				return
			}

			RenderSuccess(c, "pay murobahah request been created")
			return
		}

		//get all murobahah list from db
		murobahahs, err := DB.FindMemMurobahah(uid)
		if err != nil {
			RenderError(c, err)
			return
		}

		services.RenderPages(c, HTMLFILENAME.PayMurobahah(), murobahahs)
	}
}
