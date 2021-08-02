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
	r.Static("/public", "./public")

	r.GET(route.Home(), controller.Home)
	r.GET(route.EditCOA(), controller.EditCOA)
	r.GET(route.EditTransaction(), controller.EditTransaction)
	r.GET(route.NewCOA(), controller.NewCOA)
	r.GET(route.NewTransaction(), controller.NewTransaction)
	r.GET(route.ManageUser(), controller.ManageUser)
	r.GET(route.Summary(), controller.Summary)
	r.GET(route.FullJournal(), controller.FullJournal)

	r.Run()

}
