package chapter03

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Chap04Controller struct {
	beego.Controller
}

// 自定义模板函数
func SubStrAndRep(str string, start_num, end_num int) string {
	fmt.Println(str)
	fmt.Println(start_num)
	fmt.Println(end_num)

	bt := []rune(str)
	if len(str) > end_num { //正常切

		return string(bt[start_num:end_num]) + "..."
	} else {
		return str
	}

}

func (c *Chap04Controller) Get() {
	c.TplName = "chapter03/chap04.html"

}
