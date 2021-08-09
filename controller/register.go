package controller

import (
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var SECRET string

func init() {
	_ = godotenv.Load()
	SECRET = os.Getenv("SECRET")
}

func Register(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			memid := c.PostForm(konstanta.QueryMemID)
			username := c.PostForm(konstanta.QueryUsername)
			pass := c.PostForm(konstanta.QueryPass)
			role := c.PostForm(konstanta.QueryRole)

			//hashing pass before insert into db
			hashedPassbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
			if err != nil {
				RenderError(c, err)
				return
			}

			_, err = DB.InsertUser(memid, username, string(hashedPassbyte), role)
			if err != nil {
				RenderError(c, err)
				return
			}
			RenderSuccess(c, "user registered")
			return
		}
		services.RenderPages(c, HTMLFILENAME.Register(), nil)
	}
}
