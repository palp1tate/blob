package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id         int       `orm:"pk;auto"`
	UserName   string    `orm:"description(用户名);index;unique"`
	Password   string    `orm:"description(密码)"`
	IsAdmin    int       `orm:"description(1是管理员,2是普通用户);default(2)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	Avatar     string    `orm:"description(头像);default(/static/upload/avatar.jpg)"`

	Posts []*Post `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Comment))
}
