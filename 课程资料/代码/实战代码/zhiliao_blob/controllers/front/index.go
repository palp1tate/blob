package front

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhiliao_blob/models"
)

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {

	o := orm.NewOrm()

	posts := []models.Post{}
	o.QueryTable(new(models.Post)).RelatedSel().All(&posts)
	front_user_name := i.GetSession("front_user_name")

	if front_user_name == nil {
		front_user_name = ""
	}

	i.Data["username"] = front_user_name
	i.Data["posts"] = posts
	i.TplName = "front/index.html"

}

//[
//	{
//		id:1,
//		content:"xxx",
//		children:[
//			{
//				id:2,
//				content:"xxx",
//			},
//			{
//				id:3,
//				content:"xxx",
//			}
//		]
//
//	}
//]

func (i *IndexController) PostDetail() {

	id, _ := i.GetInt("id")

	o := orm.NewOrm()

	post := models.Post{}
	qs := o.QueryTable(new(models.Post)).Filter("id", id)
	qs.RelatedSel().One(&post)

	// 阅读数+1
	qs.Update(orm.Params{"read_num": post.ReadNum + 1})

	front_user_name := i.GetSession("front_user_name")

	if front_user_name == nil {
		front_user_name = ""
	}

	comments := []models.Comment{}
	o.QueryTable(new(models.Comment)).Filter("post_id", id).Filter("p_id", 0).RelatedSel().All(&comments)

	comment_trees := []models.CommentTree{}
	for _, comment := range comments {
		pid := comment.Id
		comment_tree := models.CommentTree{
			Id:         comment.Id,
			Content:    comment.Content,
			Author:     comment.Author,
			CreateTime: comment.CreateTime,
			Children:   []*models.CommentTree{},
		}
		GetChild(pid, &comment_tree)
		comment_trees = append(comment_trees, comment_tree)

	}

	i.Data["username"] = front_user_name
	i.Data["post"] = post
	i.Data["comment_trees"] = comment_trees
	i.TplName = "front/detail.html"

}

// 递归
func GetChild(pid int, comment_tree *models.CommentTree) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Comment))
	commments := []models.Comment{}
	_, err := qs.Filter("p_id", pid).RelatedSel().All(&commments)
	if err != nil {
		return
	}

	// 查询二级及以下的多级评论
	for i := 0; i < len(commments); i++ {
		pid := commments[i].Id
		child := models.CommentTree{Id: commments[i].Id, Content: commments[i].Content, Author: commments[i].Author, CreateTime: commments[i].CreateTime, Children: []*models.CommentTree{}}
		comment_tree.Children = append(comment_tree.Children, &child)
		GetChild(pid, &child)
	}

	return
}
