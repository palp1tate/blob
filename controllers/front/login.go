package front

import (
	"blob/models"
	"blob/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type FrontLoginController struct {
	web.Controller
}

func (r *FrontLoginController) Get() {
	r.TplName = "front/login.html"

}

func (r *FrontLoginController) Post() {

	username := r.GetString("username") // 特殊字符过滤。?/
	password := r.GetString("password") // 长度校验

	md5Pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5Pwd).Filter("IsAdmin", 2).Exist()

	if exist {

		r.SetSession("front_user_name", username)
		r.Redirect(web.URLFor("MainController.Get"), 302)
	} else {
		r.Redirect(web.URLFor("FrontLoginController.Get"), 302)
	}

}
