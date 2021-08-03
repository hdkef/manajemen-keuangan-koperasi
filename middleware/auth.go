package middleware

import (
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var route = konstanta.GetRoute()

func Auth(c *gin.Context) {
	//if bearer token valid, redirect
	claims, err := services.ValidateTokenFromCookies(c)
	if err == nil {
		//go to destination
		c.Set(konstanta.Claims, claims)
		c.Next()
		return
	}
	//if URL is not /login next
	if c.Request.URL.Path != route.Login() {
		c.Redirect(http.StatusTemporaryRedirect, route.Login())
		return
	}
	//if URL is /login next
	c.Next()
}
