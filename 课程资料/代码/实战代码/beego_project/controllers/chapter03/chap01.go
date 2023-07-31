package chapter03

import "github.com/astaxie/beego"

type Chap01Controller struct {
	beego.Controller
}

func (c *Chap01Controller) Get() {

	c.Data["name"] = "hallen"

	c.Data["arr"] = []int{1}

	c.TplName = "chapter03/chap01.html"

}
