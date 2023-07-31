package front

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhiliao_blob/models"
	"zhiliao_blob/utils"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "front/register.html"

}

func (r *RegisterController) Post() {

	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")

	if password != repassword {
		r.Ctx.WriteString("两次密码不一致")
	}

	md5_password := utils.GetMd5(password)

	o := orm.NewOrm()

	user := models.User{
		UserName: username,
		Password: md5_password,
		IsAdmin:  2,
		Cover:    "static/upload/bq3.png",
	}
	_, err := o.Insert(&user)

	if err != nil {
		r.Ctx.WriteString("用户名已被使用")
	}

	r.Redirect("/login", 302)

}
