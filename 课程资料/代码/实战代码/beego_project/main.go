package main

import (
	"beego_project/controllers/chapter03"
	"beego_project/controllers/chapter04"
	_ "beego_project/models/chapter05"
	_ "beego_project/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/logs"
)

func init() {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, database)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", datasource)

	fmt.Println("连接数据")

	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)

	if err != nil {
		panic(err)
	}

	//orm.Debug = true

}

func main() {
	//beego.SetStaticPath("/front","front")
	beego.AddFuncMap("substr_rep", chapter03.SubStrAndRep)

	// 过滤器
	beego.InsertFilter("/cms/*", beego.BeforeRouter, chapter04.MyFilter)

	orm.RunCommand()

	// logs
	//logs.SetLogger(logs.AdapterConsole)
	//logs.SetLogger(logs.AdapterFile,`{"filename":"logs/beego_project.log"}`)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/beego_project.log","separate":["error","info"]}`)
	beego.Run()
}
