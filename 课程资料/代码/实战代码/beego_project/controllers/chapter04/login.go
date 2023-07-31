package chapter04

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {

	// 假设登录成功
	// 设置session
	l.SetSession("user_name", "hallen")
	user := l.GetSession("user_name")
	fmt.Println("+++++++++获取session")
	fmt.Println(user)

	fmt.Println("+++++++++删除session")
	l.DelSession("user_name")

	user1 := l.GetSession("user_name")
	fmt.Println("+++++++++再次获取session")
	fmt.Println(user1)

	l.Ctx.WriteString("设置session成功")

}
