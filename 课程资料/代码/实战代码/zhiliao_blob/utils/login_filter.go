package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func CmsLoginFilter(ctx *context.Context) {

	cms_user_name := ctx.Input.Session("cms_user_name")

	if cms_user_name == nil {
		ctx.Redirect(302, beego.URLFor("LoginController.Get"))
	}

}

func FrontLoginFilter(ctx *context.Context) {

	front_user_name := ctx.Input.Session("front_user_name")

	if front_user_name == nil {
		ctx.Redirect(302, beego.URLFor("FrontLoginController.Get"))
	}

}
