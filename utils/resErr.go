package utils

import "github.com/gin-gonic/gin"

func ResErr(c *gin.Context, errcode int, err *error) {
	c.Writer.WriteHeader(errcode)
	c.Writer.Write([]byte((*err).Error()))
}
