package front

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"

	"blob/models"
)

type CommentController struct {
	web.Controller
}

func (c *CommentController) Post() {
	postId, err := c.GetInt("post_id")

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	post := models.Post{}
	o.QueryTable(new(models.Post)).Filter("id", postId).One(&post)

	content := c.GetString("content")
	userName := c.GetSession("front_user_name")

	if userName == nil {
		c.Data["json"] = map[string]interface{}{"code": 401, "msg": "未登录"}
		c.ServeJSON()
		return
	}

	user := models.User{}
	o.QueryTable(new(models.User)).Filter("user_name", userName).One(&user)

	pid, err1 := c.GetInt("pid")

	if err1 != nil {
		pid = 0
	}

	comment := models.Comment{
		Post:    &post,
		Content: content,
		PId:     pid,
		Author:  &user,
	}

	_, err3 := o.Insert(&comment)

	if err3 != nil {
		c.Data["json"] = map[string]interface{}{"code": 500, "msg": "评论出错，请重试"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code": 200, "msg": "评论成功"}
	c.ServeJSON()

}
