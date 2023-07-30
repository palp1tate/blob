package cms

import "github.com/beego/beego/v2/server/web"

type IndexController struct {
	web.Controller
}

func (c *IndexController) Get() {
	c.TplName = "cms/index.html"
}

func (c *IndexController) Welcome() {
	c.TplName = "cms/welcome.html"
}
