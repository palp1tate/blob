package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "zhiliao_blob/models"
	_ "zhiliao_blob/routers"
	"zhiliao_blob/utils"
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
	beego.InsertFilter("/cms/main/*", beego.BeforeRouter, utils.CmsLoginFilter)
	//beego.InsertFilter("/comment",beego.BeforeRouter,utils.FrontLoginFilter)
	orm.RunCommand()
	beego.Run()
}
