package controller

import (
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

func DepositReq(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			//get type and amount of deposit req
			user, exist := c.Get(konstanta.Claims)
			if !exist {
				RenderError(c, errors.New("claims not exist"))
				return
			}
			type_ := c.PostForm(konstanta.QueryType)
			amount, err := strconv.ParseFloat(c.PostForm(konstanta.QueryAmount), 64)
			if err != nil {
				RenderError(c, err)
				return
			}
			info := c.PostForm(konstanta.QueryInfo)

			_, err = DB.InsertMemReq(user.(models.User).ID, type_, amount, driver.MemReqOption{
				DueDate: time.Now(),
				Info:    info,
			})
			if err != nil {
				RenderError(c, err)
				return
			}
			RenderSuccess(c, "deposit request created")
			return
		}
		services.RenderPages(c, HTMLFILENAME.DepositReq(), nil)
	}
}
