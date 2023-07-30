package cms

import (
	"blob/models"
	"blob/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

var defaultCoverPath = "/static/upload/cover.jpg"

type PostController struct {
	web.Controller
}

func (c *PostController) Get() {
	o := orm.NewOrm()
	var posts []models.Post

	qs := o.QueryTable(new(models.Post))
	qs.RelatedSel().All(&posts)

	count, _ := qs.Count()
	//分页
	currentPage, _ := c.GetInt("page", 1)
	pageSize := 10

	//总页数
	totalPages := utils.GetPageNum(count, pageSize)
	//前后页码
	aroundPage := 4
	leftPages, rightPages, leftHasMore, rightHasMore := utils.GetPaginationData(totalPages, currentPage, aroundPage)

	hasPrePage, hasNextPage, prePage, nextPage := utils.HasNext(currentPage, totalPages)

	c.Data["count"] = count
	c.Data["leftPages"] = leftPages
	c.Data["rightPages"] = rightPages
	c.Data["leftHasMore"] = leftHasMore
	c.Data["rightHasMore"] = rightHasMore
	c.Data["currentPage"] = currentPage
	c.Data["hasPrePage"] = hasPrePage
	c.Data["hasNextPage"] = hasNextPage
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["pageNum"] = totalPages
	c.Data["posts"] = posts
	c.TplName = "cms/post-list.html"
}

func (c *PostController) PostAdd() {
	c.TplName = "cms/post-add.html"
}

func (c *PostController) DoAdd() {

	title := c.GetString("title")
	desc := c.GetString("desc")
	content := c.GetString("content")

	path, _, err := CoverUpload(c, defaultCoverPath)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 400, "msg": err.Error()}
		c.ServeJSON()
	}

	o := orm.NewOrm()
	author := c.GetSession("cmsUsername")
	user := models.User{}
	o.QueryTable(new(models.User)).Filter("user_name", author).One(&user)

	post := models.Post{
		Title:   title,
		Desc:    desc,
		Content: content,
		Cover:   path,
		Author:  &user,
	}

	if _, err := o.Insert(&post); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 400, "msg": "添加失败"}
	}

	c.ServeJSON()

}

func (c *PostController) PostDelete() {

	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()
	_, err2 := o.QueryTable(new(models.Post)).Filter("id", id).Delete()

	if err2 != nil {
		c.Ctx.WriteString("删除错误")
	}

	c.Redirect(web.URLFor("PostController.Get"), 302)

}

func (c *PostController) PostEdit() {

	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()

	post := models.Post{}
	o.QueryTable(new(models.Post)).Filter("id", id).One(&post)
	c.Data["post"] = post
	c.TplName = "cms/post-edit.html"

}

func (c *PostController) DoEdit() {

	o := orm.NewOrm()

	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 400, "msg": "id参数错误"}
		c.ServeJSON()
	}

	qs := o.QueryTable(new(models.Post)).Filter("id", id)

	title := c.GetString("title")
	desc := c.GetString("desc")
	content := c.GetString("content")

	path, ok, err := CoverUpload(c, defaultCoverPath)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 400, "msg": err.Error()}
		c.ServeJSON()
	}

	params := orm.Params{
		"title":   title,
		"desc":    desc,
		"content": content,
	}
	if ok {
		params["cover"] = path
	}

	if _, err := qs.Update(params); err != nil {
		c.Data["json"] = map[string]interface{}{"code": 400, "msg": "更新失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
	}
	c.ServeJSON()
}

func CoverUpload(c *PostController, defaultPath string) (string, bool, error) {
	f, h, err := c.GetFile("cover")
	if err == nil {
		defer f.Close()

		// 生成时间戳，防止重名
		timeUnix := time.Now().Unix()
		timeStr := strconv.FormatInt(timeUnix, 10)

		path := "static/upload/" + timeStr + h.Filename

		// 保存获取到的文件
		if err := c.SaveToFile("cover", path); err != nil {
			return defaultPath, false, err
		}

		return "/" + path, true, nil
	}

	return defaultPath, false, nil
}
