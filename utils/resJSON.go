package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResJSON(c *gin.Context, data interface{}) {
	err := json.NewEncoder(c.Writer).Encode(data)
	if err != nil {
		ResErr(c, http.StatusInternalServerError, &err)
		return
	}
}
