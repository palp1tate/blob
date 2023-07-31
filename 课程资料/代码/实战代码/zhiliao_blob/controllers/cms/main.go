package cms

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.TplName = "cms/index.html"

}

func (m *MainController) Welcome() {
	m.TplName = "cms/welcome.html"

}
