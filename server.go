package main

import (
	"manajemen-keuangan-koperasi/controller"
	"manajemen-keuangan-koperasi/driver"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {

	db := driver.DBConn()
	defer db.DB.Close()

	route := konstanta.GetRoute()

	r := gin.Default()
	r.LoadHTMLGlob("public/pages/*")
	r.Static("/public", "./public")

	r.Use(middleware.Auth)
	r.Use(middleware.DB(db))

	r.GET(route.Home(), controller.Home)
	r.GET(route.EditCOA(), controller.EditCOA)
	r.GET(route.EditTransaction(), controller.EditTransaction)
	r.GET(route.NewCOA(), controller.NewCOA)
	r.GET(route.NewTransaction(), controller.NewTransaction)
	r.GET(route.ManageUser(), controller.ManageUser)
	r.GET(route.Summary(), controller.Summary)
	r.GET(route.FullReport(), controller.FullReport)
	r.GET(route.Admin(), controller.Admin)
	r.GET(route.Login(), controller.Login)
	r.GET(route.Member(), controller.Member)
	r.GET(route.MemRequest(), controller.MemRequest)
	r.GET(route.LoanReq(), controller.LoanReq)
	r.GET(route.MemInspect(), controller.MemInspect)
	r.GET(route.DepositReq(), controller.DepositReq)
	r.GET(route.WithdrawReq(), controller.WithdrawReq)
	r.GET(route.FindUser(), controller.FindUser)
	r.GET(route.UserSetting(), controller.UserSetting)
	r.GET(route.Register(), controller.Register)

	r.POST(route.Login(), controller.Login)

	r.Run()

}
