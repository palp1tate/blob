package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (c *MainController) Get() {

	// 渲染字符串到浏览器:不会渲染到模板文件
	c.Ctx.WriteString("hello beego")

	// 结构体渲染
	//var user User
	//user.Id = 1
	//user.Name = "hallen"
	//user.Age = 18

	user := User{
		Id:   2,
		Name: "hallen",
		Age:  18,
	}

	// 数组渲染
	arr := [5]int{1, 2, 3, 4, 5}

	// 结构体+数组

	users := [3]User{
		{
			Id:   1,
			Name: "hallen1",
			Age:  18,
		},
		{
			Id:   2,
			Name: "hallen2",
			Age:  19,
		},
		{
			Id:   3,
			Name: "hallen3",
			Age:  20,
		},
	}

	// map渲染
	//map_data := map[string]string{
	//	"name":"hallen",
	//	"age":"18",
	//}
	map_data := map[string]interface{}{
		"name": "hallen",
		"age":  18,
	}
	// map+结构体渲染
	map_struct := map[string]User{
		"user1": {
			Id:   1,
			Name: "hallen1",
			Age:  18,
		},
		"user2": {
			Id:   2,
			Name: "hallen2",
			Age:  19,
		},
	}

	// 切片渲染
	slice := []int{1, 2, 3, 4, 5, 5, 7, 524, 13, 1}
	c.Data["user"] = user
	c.Data["arr"] = arr
	c.Data["users"] = users
	c.Data["map_data"] = map_data
	c.Data["map_struct"] = map_struct
	c.Data["slice"] = slice

	c.TplName = "index.html"
}
