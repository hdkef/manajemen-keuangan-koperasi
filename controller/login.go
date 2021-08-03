package controller

import (
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"
	"manajemen-keuangan-koperasi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		//TOBE

		//retrieve user from db and compare the pass
		usr := models.User{
			ID:       "1",
			Username: "dk",
		}

		//create token and save
		tokenString, err := services.GenerateTokenStringFromUserStruct(&usr)
		if err != nil {
			utils.ResErr(c, http.StatusInternalServerError, &err)
			return
		}
		services.SaveTokenCookie(c, &tokenString)

		//redirect
	}
	err := services.ValidateTokenFromCookies(c)
	if err == nil {
		//redirect
		c.Redirect(http.StatusTemporaryRedirect, route.Admin())
		return
	}
	services.RenderPages(c, HTMLFILENAME.Login(), nil)
}
