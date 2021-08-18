package controller

import (
	"context"
	"fmt"
	"log"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MurobahahDec(DB *driver.DBDriver) func(c *gin.Context) {
	return func(c *gin.Context) {
		//getting all values
		id, err := strconv.ParseFloat(c.PostForm(konstanta.QueryID), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		agentid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUIDAgent), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		buyerid, err := strconv.ParseFloat(c.PostForm(konstanta.QueryUIDBuyer), 64)
		if err != nil {
			RenderError(c, err)
			return
		}
		doc := c.PostForm(konstanta.QueryDoc)

		//Begin transaction
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		tx, err := DB.DB.BeginTx(ctx, nil)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//send info to agent & buyer (bulk insert) that request been declined
		info := fmt.Sprintf("murobahah request with id of %v been declined", id)

		_, err = DB.InsertBatchAllInfoTx(tx, []float64{agentid, buyerid}, info)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//delete req from murobahah req
		_, err = DB.DeleteMemReqMurobahahTx(tx, id)
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//commit
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			RenderError(c, err)
			return
		}

		//delete document that has been uploaded
		err = services.RemoveFile(doc)
		if err != nil {
			log.Println(err)
		}

		//Render success declined
		RenderSuccess(c, "murobahah been declined")

	}
}
