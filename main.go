package main

import (
	"api-sample/models"
	_ "api-sample/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	mysqlHost := beego.AppConfig.String("mysqlHost")
	mysqldb := beego.AppConfig.String("mysqldb")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", mysqlUser, mysqlPass, mysqlHost, mysqldb)
	orm.RegisterDataBase("default", "mysql", datasource, 30)

	orm.RegisterModel(new(models.User))

	//sync初期化時に同期をとる
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// beego Cors resoluve
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost*"},
		AllowMethods:     []string{"GET", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
