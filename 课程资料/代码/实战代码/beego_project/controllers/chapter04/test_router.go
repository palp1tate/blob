package chapter04

import "github.com/astaxie/beego"

type RouterController struct {
	beego.Controller
}

func (r *RouterController) List() {

	r.Data["name"] = "hallen"
	r.TplName = "chapter04/test_router.html"

}
