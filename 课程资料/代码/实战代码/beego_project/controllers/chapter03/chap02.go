package chapter03

import "github.com/astaxie/beego"

type Chap02Controller struct {
	beego.Controller
}

func (c *Chap02Controller) Get() {
	c.Data["x"] = ""
	c.Data["y"] = "y"
	c.Data["z"] = "z"
	c.Data["arr"] = []int{3, 4, 5, 6}
	c.Data["is_ok"] = true
	c.Data["age"] = 18
	c.Data["map_data"] = map[string]interface{}{"name": "hallen", "age": 19}
	c.TplName = "chapter03/chap02.html"

}
