package utils

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func CmsLoginFilter(ctx *context.Context) {
	cmsUserName := ctx.Input.Session("cmsUsername")
	if cmsUserName == nil {
		ctx.Redirect(302, web.URLFor("LoginController.Get"))
	}

}
