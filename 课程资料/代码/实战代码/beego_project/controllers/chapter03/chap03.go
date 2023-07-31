package chapter03

import (
	"github.com/astaxie/beego"
	"time"
)

type Chap03Controller struct {
	beego.Controller
}

func (c *Chap03Controller) Get() {
	c.Data["date"] = time.Now()

	c.Data["map_data"] = map[string]interface{}{"name": "hallen", "age": 19}
	c.TplName = "chapter03/chap03.html"

}
