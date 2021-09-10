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

	redis := driver.RedisConn()
	defer redis.C.Close()

	route := konstanta.GetRoute()

	r := gin.Default()
	r.LoadHTMLGlob("public/pages/*")
	r.Static("/public", "./public")

	//use middleware Auth
	r.Use(middleware.Auth)

	r.GET(route.Home(), controller.Home(db))
	r.GET(route.EditCOA(), controller.EditCOA(db))
	r.GET(route.EditTransaction(), controller.EditTransaction(db))
	r.GET(route.NewCOA(), controller.NewCOA(db))
	r.GET(route.NewTransaction(), controller.NewTransaction(db))
	r.GET(route.ManageUser(), controller.ManageUser(db, redis))
	r.GET(route.Summary(), controller.Summary(db))
	r.GET(route.FullReport(), controller.FullReport(db))
	r.GET(route.Admin(), controller.Admin(db))
	r.GET(route.Login(), controller.Login(db))
	r.GET(route.Member(), controller.Member(db, redis))
	r.GET(route.MemRequest(), controller.MemRequest(db))
	r.GET(route.MemInspect(), controller.MemInspect(db))
	r.GET(route.DepositReq(), controller.DepositReq(db))
	r.GET(route.WithdrawReq(), controller.WithdrawReq(db))
	r.GET(route.FindUser(), controller.FindUser(db))
	r.GET(route.UserSetting(), controller.UserSetting(db))
	r.GET(route.Register(), controller.Register(db))
	r.GET(route.AdmInput(), controller.AdmInput(db))
	r.GET(route.MurobahahReq(), controller.MurobahahReq(db))
	r.GET(route.PayMurobahah(), controller.PayMurobahah(db))

	r.POST(route.Login(), controller.Login(db))
	r.POST(route.Register(), controller.Register(db))
	r.POST(route.DepositReq(), controller.DepositReq(db))
	r.POST(route.WithdrawReq(), controller.WithdrawReq(db))
	r.POST(route.MemRequestAcc(), controller.MemRequestAcc(db))
	r.POST(route.MemRequestDec(), controller.MemRequestDec(db))
	r.POST(route.AdmInput(), controller.AdmInput(db))
	r.POST(route.MurobahahReq(), controller.MurobahahReq(db))
	r.POST(route.MurobahahAcc(), controller.MurobahahAcc(db))
	r.POST(route.MurobahahDec(), controller.MurobahahDec(db))
	r.POST(route.ManageUser(), controller.ManageUser(db, redis))
	r.POST(route.PayMurobahah(), controller.PayMurobahah(db))
	r.POST(route.PayMurobahahAcc(), controller.PayMurobahahAcc(db))

	r.Any(route.DelCacheMember(), controller.DelCacheMember(redis))
	r.Any(route.Logout(), controller.Logout(redis))

	r.Run()

}
