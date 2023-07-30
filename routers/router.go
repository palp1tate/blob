package routers

import (
	"blob/controllers/cms"
	"blob/controllers/front"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/cms", &cms.LoginController{})
	web.Router("/cms/index", &cms.IndexController{})
	web.Router("/cms/index/welcome", &cms.IndexController{}, "get:Welcome")
	web.Router("/cms/index/post", &cms.PostController{})
	web.Router("/cms/index/post_add", &cms.PostController{}, "get:PostAdd")
	web.Router("/cms/index/post_do_add", &cms.PostController{}, "post:DoAdd")
	web.Router("/cms/index/post_delete", &cms.PostController{}, "get:PostDelete")
	web.Router("/cms/index/post_edit", &cms.PostController{}, "get:PostEdit")
	web.Router("/cms/index/post_do_edit", &cms.PostController{}, "post:DoEdit")

	// 前端
	web.Router("/", &front.MainController{})
	web.Router("/detail", &front.MainController{}, "get:PostDetail")
	web.Router("/register", &front.RegisterController{})
	web.Router("/login", &front.FrontLoginController{})
	web.Router("/comment", &front.CommentController{})
}
