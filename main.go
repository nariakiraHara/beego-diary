package main

import (
	_ "api-sample/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	mysqlHost := beego.AppConfig.String("mysqlHost")
	mysqldb := beego.AppConfig.String("mysqldb")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", mysqlUser, mysqlPass, mysqlHost, mysqldb)
	orm.RegisterDataBase("default", "mysql", datasource, 30)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
