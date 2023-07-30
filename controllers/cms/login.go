package cms

import (
	"blob/models"
	"blob/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

func (c *LoginController) Get() {
	c.TplName = "cms/login.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	md5Pwd := utils.GetMd5(password)
	o := orm.NewOrm()
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 400, "msg": "用户名或密码不能为空"}
		c.ServeJSON()
		return
	}
	user := models.User{
		UserName: username,
		Password: md5Pwd,
		IsAdmin:  1,
	}
	if err := o.Read(&user, "UserName", "Password", "IsAdmin"); err == nil {
		c.SetSession("cmsUsername", username)
		c.Redirect(web.URLFor("IndexController.Get"), 302)
		return
	} else {
		c.Redirect(web.URLFor("LoginController.Get"), 302)
	}

}
