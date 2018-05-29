package main

import (
	_ "api/routers"
	"fmt"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/astaxie/beego"
)

func init() {
	DbConnection()
}

func DbConnection() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_name := beego.AppConfig.String("db_name")
	db_username := beego.AppConfig.String("db_username")
	db_password := beego.AppConfig.String("db_password")
	db_type := beego.AppConfig.String("db_type")
	orm.RegisterDriver(db_type, orm.DRMySQL)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_username, db_password, db_host, db_port, db_name)
	orm.RegisterDataBase("default", db_type, dns)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.FinishRouter, UserFilter) //开启过滤器
	beego.InsertFilter("/*", beego.BeforeExec, testFilter)   //开启过滤器
	beego.Run()
}

var UserFilter = func(ctx *context.Context) { //校验用户的合法性
	//这里就直接调用 官方的api 只不过被beego封装了一层 哦哦 我研究研究 恩 我也忘了具体怎么拦截退出  应该是直接回调然后 return 直接结束
	//官方文档应该有 你起看看吧 恩恩 等我研究好gin 又可以抱大腿了
}

var testFilter = func(ctx *context.Context) {

	fmt.Println(ctx.Input.URL())

	return
}
