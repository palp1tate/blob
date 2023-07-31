package chapter02

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	u.TplName = "chapter02/upload.html"

}

func (u *UploadController) Post() {
	// 获取上传的文件
	f, h, _ := u.GetFile("file")

	defer f.Close()
	fmt.Println(h.Filename)

	// 保存上传的文件

	time_unix := time.Now().Unix()
	fmt.Println(time_unix)

	file_path := fmt.Sprintf("%d_%s", time_unix, h.Filename)
	u.SaveToFile("file", "upload/"+file_path)
	//u.Ctx.WriteString("上传成功")

	u.Data["json"] = map[string]interface{}{"code": 200, "msg": "上传成功"}
	u.ServeJSON()

}
