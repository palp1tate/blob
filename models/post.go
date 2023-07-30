package models

import (
	"time"
)

//type Tag struct {
//	Id int `orm:"pk;auto"`
//	Name string `orm:"description(标签名称)"`
//
//}

type Post struct {
	Id         int        `orm:"pk;auto"`
	Title      string     `orm:"description(帖子标题)"`
	Desc       string     `orm:"description(帖子简介)"`
	Content    string     `orm:"size(4000);description(帖子内容)"`
	Cover      string     `orm:"description(帖子封面);default(/static/upload/cover.jpg)"`
	ReadNum    int        `orm:"description(阅读量);default(0)"`
	StarNum    int        `orm:"description(点赞数);default(0)"`
	Author     *User      `orm:"rel(fk);description(作者)"`
	Comments   []*Comment `orm:"reverse(many);description(评论)"`
	CreateTime time.Time  `orm:"auto_now_add;type(datetime);description(创建时间)"`
}

func (p *Post) TableName() string {
	return "sys_post"
}
