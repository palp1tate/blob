package front

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhiliao_blob/models"
	"zhiliao_blob/utils"
)

type FrontLoginController struct {
	beego.Controller
}

func (l *FrontLoginController) Get() {
	l.TplName = "front/login.html"

}

func (l *FrontLoginController) Post() {

	username := l.GetString("username") // 特殊字符过滤。?/
	password := l.GetString("password") // 长度校验

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5_pwd).Exist()

	if exist {

		l.SetSession("front_user_name", username)
		fmt.Println("登录成功")
		l.Redirect(beego.URLFor("IndexController.Get"), 302)
	} else {
		l.Redirect(beego.URLFor("FrontLoginController.Get"), 302)
	}

}
