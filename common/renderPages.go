package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderPages(c *gin.Context, fname string, data interface{}) {
	c.HTML(http.StatusOK, fname, data)
}
