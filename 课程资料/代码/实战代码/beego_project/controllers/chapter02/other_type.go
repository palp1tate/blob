package chapter02

import "github.com/astaxie/beego"

type OtherTypeController struct {
	beego.Controller
}

type Teacher struct {
	Id   int
	Name string
	Age  int
}

func (o *OtherTypeController) Get() {

	teacher := Teacher{Id: 1, Name: "hallen", Age: 18}

	// json格式
	//o.Data["json"] = teacher
	//o.ServeJSON()

	// xml
	//o.Data["xml"] = teacher
	//o.ServeXML()

	// yaml
	o.Data["yaml"] = teacher
	o.ServeYAML()
}
