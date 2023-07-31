package chapter04

import (
	"github.com/astaxie/beego/context"
)

func MyFilter(ctx *context.Context) {

	user_name := ctx.Input.Session("user_name")

	if user_name == nil {
		ctx.WriteString("没有登录")
	}

}
