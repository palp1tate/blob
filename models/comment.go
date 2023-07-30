package models

import "time"

type CommentTree struct {
	Id         int            `orm:"pk;auto"`
	Content    string         `orm:"size(4000);description(评论内容)"`
	Author     *User          `orm:"rel(fk);description(评论人)"`
	CreateTime time.Time      `orm:"auto_now_add;type(datetime);description(创建时间)"`
	Children   []*CommentTree `orm:"reverse(many);description(子评论)"`
}

type Comment struct {
	Id      int    `orm:"pk;auto"`
	Content string `orm:"size(4000);description(评论内容)"`

	Post *Post `orm:"rel(fk);description(帖子外键)"`
	// 评论的父级id
	PId        int       `orm:"description(父级评论);default(0)"`
	Author     *User     `orm:"rel(fk);description(评论人)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
}

//1 xx 0
//2 qq 1

func (c *Comment) TableName() string {
	return "sys_post_comment"

}
