package middleware

import (
	"errors"
	"manajemen-keuangan-koperasi/controller"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var route = konstanta.GetRoute()

func Auth(c *gin.Context) {
	path := c.Request.URL.Path
	//if bearer token valid, redirect
	claims, err := services.ValidateTokenFromCookies(c)
	//if there is valid token
	if err == nil {
		//send claims to context
		c.Set(konstanta.Claims, claims)
		//if must be admin
		if path == route.Admin() || path == route.EditCOA() || path == route.EditTransaction() || path == route.FullReport() || path == route.NewCOA() || path == route.NewTransaction() || path == route.Summary() || path == route.MemRequest() || path == route.MemInspect() || path == route.FindUser() || path == route.Register() || path == route.MemRequestAcc() || path == route.MemRequestDec() {
			mustAdmin(c, &claims)
			return
		}
		if path == route.ManageUser() {
			mustAdminSuper(c, &claims)
			return
		}
		//if token is valid and it is authenticated, redirect to member
		if path == route.Login() {
			c.Redirect(http.StatusTemporaryRedirect, route.Member())
			return
		}
		//go to destination
		c.Next()
		return
	}
	//if no valid token and URL is not /login next
	if path != route.Login() {
		c.Redirect(http.StatusTemporaryRedirect, route.Login())
		return
	}
	//if no valid token and URL is /login next
	c.Next()
}

func mustAdmin(c *gin.Context, claims *models.User) {
	if claims.Role == konstanta.RoleADMINInput || claims.Role == konstanta.RoleADMINSuper {
		c.Next()
		return
	}
	c.Abort()
	controller.RenderError(c, errors.New("USER IS NOT ADMIN"))
}

func mustAdminSuper(c *gin.Context, claims *models.User) {
	if claims.Role == konstanta.RoleADMINSuper {
		c.Next()
		return
	}
	c.Abort()
	controller.RenderError(c, errors.New("USER IS NOT ADMIN Super"))
}
