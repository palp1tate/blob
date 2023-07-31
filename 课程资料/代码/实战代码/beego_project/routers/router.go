package routers

import (
	"beego_project/controllers"
	"beego_project/controllers/chapter02"
	"beego_project/controllers/chapter03"
	"beego_project/controllers/chapter04"
	"beego_project/controllers/chapter05"
	"beego_project/controllers/chapter06"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
	//beego.Router("/user/?:id:int",&controllers.UserController{})
	beego.Router("/user", &chapter02.UserController{})
	beego.Router("/add_user", &chapter02.UserController{})
	beego.Router("/add_user_ajax", &chapter02.UserController{})

	beego.Router("/upload", &chapter02.UploadController{})

	beego.Router("/other", &chapter02.OtherTypeController{})

	beego.Router("/chap01", &chapter03.Chap01Controller{})
	beego.Router("/chap02", &chapter03.Chap02Controller{})
	beego.Router("/chap03", &chapter03.Chap03Controller{})
	beego.Router("/chap04", &chapter03.Chap04Controller{})

	beego.Router("/test_conf", &chapter04.ConfController{})

	// 路由
	beego.Router("/test_router", &chapter04.RouterController{}, "get:List")

	// 测试过滤器的url
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &chapter04.LoginController{})

	// 需要登录的
	beego.Router("/cms/user/test1", &chapter04.ConfController{})
	beego.Router("/cms/test2", &chapter04.ConfController{})
	beego.Router("/cms/test3", &chapter04.ConfController{})
	beego.Router("/cms/test4", &chapter04.ConfController{})

	beego.Router("/test_valid", &chapter04.ValidController{})

	beego.Router("/test_exper", &chapter05.ExperController{})
	beego.Router("/test_query_table", &chapter05.TestQueryTable{})
	beego.Router("/test_logs", &chapter06.LogsController{})

}
