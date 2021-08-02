package main

import (
	"manajemen-keuangan-koperasi/controller"
	"manajemen-keuangan-koperasi/konstanta"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {

	route := konstanta.GetRoute()

	r := gin.Default()
	r.LoadHTMLGlob("public/pages/*")

	r.GET(route.Home(), controller.Home)

	r.Run()

}
