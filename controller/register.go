package controller

import (
	"context"
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
			isagent := c.PostForm(konstanta.QueryIsAgent)
			tel := c.PostForm(konstanta.QueryTel)

			//hashing pass before insert into db
			hashedPassbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
			if err != nil {
				RenderError(c, err)
				return
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			//start transaction
			tx, err := DB.DB.BeginTx(ctx, nil)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//insert user to alluser
			res, err := DB.InsertUserTx(tx, memid, username, string(hashedPassbyte), role, isagent, tel)
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			id, err := res.LastInsertId()
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//create zero balance account
			_, err = DB.CreateZeroBalance(tx, float64(id))
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			//commit transaction
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				RenderError(c, err)
				return
			}

			RenderSuccess(c, "user registered")
			return
		}
		services.RenderPages(c, HTMLFILENAME.Register(), nil)
	}
}
