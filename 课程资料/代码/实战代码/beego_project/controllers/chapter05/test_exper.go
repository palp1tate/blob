package chapter05

import (
	"beego_project/models/chapter05"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ExperController struct {
	beego.Controller
}

func (e *ExperController) Get() {

	o := orm.NewOrm()
	//o.QueryTable("sys_user")
	qs := o.QueryTable(new(chapter05.User))
	user := chapter05.User{}
	// exact
	//qs.Filter("name","hallen").One(&user)
	// contains
	//qs.Filter("name__contains","lle").One(&user)
	// gt/gte
	//qs.Filter("age__gt",19).One(&user)
	//lt/lte
	//qs.Filter("age__lt",19).One(&user)

	// startswith
	//qs.Filter("name__startswith","ha").One(&user)
	// endswith
	//qs.Filter("name__endswith","en").One(&user)

	// in
	//qs.Filter("age__in",15,16,27,19,18).One(&user)

	// isnull ,true查询为null的
	qs.Filter("address__isnull", true).One(&user)
	fmt.Println(user)
	e.Data["user"] = user
	e.TplName = "chapter05/test_exper.html"

}
