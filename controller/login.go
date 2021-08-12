package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"
	"manajemen-keuangan-koperasi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {

			//Get username and pass from form

			username := c.PostForm(konstanta.QueryUsername)
			pass := c.PostForm(konstanta.QueryPass)

			//Get user info from db

			user, err := DB.FindOneUserByUsername(username)
			if err != nil {
				RenderError(c, err)
				return
			}

			//compare pass from form and db

			err = bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(pass))
			if err != nil {
				RenderError(c, err)
				return
			}

			//create token and save
			tokenString, err := services.GenerateTokenStringFromUserStruct(&user)
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
