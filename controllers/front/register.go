package front

import (
	"blob/models"
	"blob/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	web.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "front/register.html"

}

func (c *RegisterController) Post() {

	username := c.GetString("username")
	password := c.GetString("password")
	rePassword := c.GetString("rePassword")

	if password != rePassword {
		c.Ctx.WriteString("两次密码不一致")
		return
	}

	md5Password := utils.GetMd5(password)

	o := orm.NewOrm()
	user := models.User{
		UserName: username,
		Password: md5Password,
		IsAdmin:  2,
		Avatar:   "/static/upload/avatar.jpg",
	}
	//判断用户名是否重复
	err := o.Read(&user, "UserName")
	if err == nil {
		c.Ctx.WriteString("用户名已存在")
		return

	}
	//插入数据
	_, err = o.Insert(&user)
	if err != nil {
		c.Ctx.WriteString("注册失败")
		return
	}

	c.Redirect("/login", 302)

}
