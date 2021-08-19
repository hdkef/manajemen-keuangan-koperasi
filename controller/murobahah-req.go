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

func MurobahahReq(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			buyerid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUIDBuyer), 64)
			if err != nil {
				RenderError(c, err)
				return
			}

			duedate := c.PostForm(konstanta.QueryDueDate)

			amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
			if err != nil {
				RenderError(c, err)
				return
			}

			info := c.PostForm(konstanta.QueryInfo)

			claims, exist := c.Get(konstanta.Claims)
			if !exist {
				RenderError(c, errors.New("NO MEMID"))
				return
			}

			user := claims.(models.User)

			memid := user.MemID

			agentid := user.ID

			//first upload document
			fpath, err := services.UploadFile(c, konstanta.QueryDoc, memid, "MRBH")
			if err != nil {
				RenderError(c, err)
				return
			}

			//insert into murobahah_req table
			_, err = DB.InsertMemReqMurobahah(agentid, buyerid, duedate, amount, fpath, info)
			if err != nil {
				//handle deleting file if error
				services.RemoveFile(fpath)
				RenderError(c, err)
				return
			}

			//Success response
			RenderSuccess(c, "murobahah request created")
			return
		}
		services.RenderPages(c, HTMLFILENAME.MurobahahReq(), nil)
	}
}
