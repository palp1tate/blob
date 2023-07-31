package chapter05

import (
	"beego_project/models/chapter05"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestQueryTable struct {
	beego.Controller
}

func (t *TestQueryTable) Get() {

	o := orm.NewOrm()
	user := chapter05.User{}
	users := []chapter05.User{}
	qs := o.QueryTable(new(chapter05.User))
	// one
	//qs.One(&user)

	// filter
	qs.Filter("name", "hallenx").One(&user)

	// all
	//qs.All(&users)

	// exclude
	//qs.Exclude("name","hallen").All(&users)

	// limit offset
	//qs.Limit(5).Offset(1).All(&users)

	// GroupBy
	//qs.GroupBy("age").All(&users)

	// orderBy
	//qs.OrderBy("-age").All(&users)

	// Distinct
	//qs.Distinct().All(&users,"age","address")

	// count
	//count,_ := qs.Count()
	count, _ := qs.Filter("age", 18).Count()

	// exist
	exist := qs.Filter("age", 21).Exist()

	// update
	//qs.Filter("age",18).Update(orm.Params{
	//	"address":"rrrr",
	//})

	// delete
	//qs.Filter("age",11).Delete()

	// insert
	user2 := chapter05.User{}
	user2.Name = "zhiliao7"
	profile := chapter05.Profile{}
	o.QueryTable(new(chapter05.Profile)).Filter("id", 4).One(&profile)
	fmt.Println("==============")
	fmt.Println(profile)
	user2.Profile = &profile
	_, err := o.Insert(&user2)
	fmt.Println(err)
	t.Data["user"] = user
	t.Data["users"] = users
	t.Data["count"] = count
	t.Data["exist"] = exist
	t.TplName = "chapter05/test_query_table.html"

}
