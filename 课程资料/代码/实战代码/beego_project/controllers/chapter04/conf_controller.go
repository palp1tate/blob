package chapter04

import "github.com/astaxie/beego"

type ConfController struct {
	beego.Controller
}

func (c *ConfController) Get() {
	httpport, _ := beego.AppConfig.Int("httpport")
	name := beego.AppConfig.String("name")

	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	c.Data["httpport"] = httpport
	c.Data["name"] = name
	c.Data["username"] = username
	c.Data["password"] = password
	c.Data["host"] = host
	c.Data["port"] = port
	c.Data["database"] = database

	c.TplName = "chapter04/test_conf.html"

}
