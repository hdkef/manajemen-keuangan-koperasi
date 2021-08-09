package controller

import (
	"errors"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"
	"manajemen-keuangan-koperasi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {

			//Get username and pass from form

			username := c.PostForm(konstanta.QueryUsername)
			pass := c.PostForm(konstanta.QueryPass)

			//TOBE
			//compare pass from form and db

			if pass != "tes" {
				RenderError(c, errors.New("login failed"))
				return
			}

			usr := models.User{
				ID:       "1",
				Username: username,
				Role:     konstanta.RoleADMINSuper,
			}

			//create token and save
			tokenString, err := services.GenerateTokenStringFromUserStruct(&usr)
			if err != nil {
				utils.ResErr(c, http.StatusInternalServerError, &err)
				return
			}
			services.SaveTokenCookie(c, &tokenString)

			c.Redirect(http.StatusFound, route.Member())
			return
		}
		services.RenderPages(c, HTMLFILENAME.Login(), nil)
	}
}
