package chapter05

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int    `orm:"pk;auto"`
	Name       string `orm:"index;unique"`
	Age        int
	Addr       string    `orm:"null;column(address)"`
	Desc       string    `orm:"size(2000)"`
	Price      float64   `orm:"digits(12);decimals(4)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	// 1启用，0停用
	Status   int        `orm:"default(1);description(状态，1启用，0停用)"`
	XXX      string     `orm:"-"`
	Profile  *Profile   `orm:"rel(one)"`
	Articles []*Article `orm:"reverse(many)"`
}

type Profile struct {
	Id     int
	IdCard string
	Cover  string
	User   *User `orm:"reverse(one)"`
}

type Article struct {
	Id      int
	Title   string
	Content string `orm:"size(2000)"`
	User    *User  `orm:"rel(fk)"`
}

type Post struct {
	Id   int
	Tile string
	Tags []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (p *Profile) TableName() string {
	return "sys_profile"

}

func (a *Article) TableName() string {
	return "sys_article"

}
func (a *Post) TableName() string {
	return "sys_post"

}
func (a *Tag) TableName() string {
	return "sys_tag"

}

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Article), new(Post), new(Tag))

	orm.Debug = true
}
