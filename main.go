package main

import (
	_ "blob/models"
	_ "blob/routers"
	"blob/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username, _ := web.AppConfig.String("username")
	password, _ := web.AppConfig.String("password")
	host, _ := web.AppConfig.String("host")
	port, _ := web.AppConfig.String("port")
	database, _ := web.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local", username, password, host, port, database)
	err := orm.RegisterDataBase("default", "mysql", datasource)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}

func main() {
	web.InsertFilter("/cms/index/*", web.BeforeRouter, utils.CmsLoginFilter)
	orm.RunCommand()
	web.Run()
}
