package front

import (
	"blob/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {

	o := orm.NewOrm()

	var posts []models.Post
	o.QueryTable(new(models.Post)).RelatedSel().All(&posts)
	frontUserName := c.GetSession("front_user_name")

	if frontUserName == nil {
		frontUserName = ""
	}

	c.Data["username"] = frontUserName
	c.Data["posts"] = posts
	c.TplName = "front/index.html"

}

func (c *MainController) PostDetail() {

	id, _ := c.GetInt("id")

	o := orm.NewOrm()

	post := models.Post{}
	qs := o.QueryTable(new(models.Post)).Filter("id", id)
	qs.RelatedSel().One(&post)

	// 阅读数+1
	qs.Update(orm.Params{"read_num": post.ReadNum + 1})

	frontUserName := c.GetSession("front_user_name")

	if frontUserName == nil {
		frontUserName = ""
	}

	var comments []models.Comment
	o.QueryTable(new(models.Comment)).Filter("post_id", id).Filter("p_id", 0).RelatedSel().All(&comments)

	var commentTrees []models.CommentTree

	for _, comment := range comments {
		pid := comment.Id
		//fmt.Println(pid)
		commentTree := models.CommentTree{
			Id:         comment.Id,
			Content:    comment.Content,
			Author:     comment.Author,
			CreateTime: comment.CreateTime,
			Children:   []*models.CommentTree{},
		}
		GetChild(pid, &commentTree)
		commentTrees = append(commentTrees, commentTree)

	}

	c.Data["username"] = frontUserName
	c.Data["post"] = post
	c.Data["comment_trees"] = commentTrees
	c.TplName = "front/detail.html"

}

// GetChild 递归
func GetChild(pid int, commentTree *models.CommentTree) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Comment))
	var comments []models.Comment
	_, err := qs.Filter("p_id", pid).RelatedSel().All(&comments)
	if err != nil {
		return
	}

	// 查询二级及以下的多级评论
	for i := 0; i < len(comments); i++ {
		pid := comments[i].Id
		child := models.CommentTree{
			Id:         comments[i].Id,
			Content:    comments[i].Content,
			Author:     comments[i].Author,
			CreateTime: comments[i].CreateTime,
			Children:   []*models.CommentTree{},
		}
		commentTree.Children = append(commentTree.Children, &child)
		GetChild(pid, &child)
	}

	return
}
