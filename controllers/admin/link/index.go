package SortController

import (
	"beego_blog_mvc/controllers"
	"beego_blog_mvc/models"

	"github.com/jaevor/go-nanoid"
)

type LinkController struct {
	controllers.Nuzn
}

// GetAll 获取友链
func (c *LinkController) GetAll() {
	ml := models.NewLink()
	linkAll := ml.GetAdminLinkList()
	c.Data["linkAll"] = linkAll
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/link/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/link/css.html"
	c.LayoutSections["LayoutJs"] = "admin/link/js.html"
}

// Add 渲染添加友链
func (c *LinkController) Add() {
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/link/add_link.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/link/css.html"
	c.LayoutSections["LayoutJs"] = "admin/link/js.html"
}

// AddPost 添加友链的post
func (c *LinkController) AddPost() {
	// 接收参数
	title := c.GetString("title")
	url := c.GetString("url")
	sort, _ := c.GetInt("sort")
	show, _ := c.GetInt("show")
	describe := c.GetString("describe")

	ml := models.NewLink()
	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	ml.Uid = uid
	ml.Title = title
	ml.Url = url
	ml.Show = show
	ml.Sort = sort
	ml.Describe = describe
	// 业务处理
	err := ml.Create()
	if err != nil {
		c.Fail("创建友链失败")
		return
	} else {
		c.Ok("创建友链成功")
		return
	}
}

// Edit 渲染 编辑友链
func (c *LinkController) Edit() {
	id := c.GetString("id")
	ml := models.NewLink()
	sort := ml.GetLinkByUid(id)
	c.Data["Result"] = sort
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/link/edit_link.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/link/css.html"
	c.LayoutSections["LayoutJs"] = "admin/link/js.html"
}

// EditPost 编辑友链的post
func (c *LinkController) EditPost() {
	// 接收参数
	id := c.GetString("id")
	title := c.GetString("title")
	url := c.GetString("url")
	sort, _ := c.GetInt("sort")
	show, _ := c.GetInt("show")
	describe := c.GetString("describe")

	ml := models.NewLink()
	// 组装数据
	ml.Title = title
	ml.Url = url
	ml.Sort = sort
	ml.Show = show
	ml.Describe = describe

	// 业务处理
	err := ml.Update(id)
	if err != nil {
		c.Fail("编辑友链失败")
		return
	} else {
		c.Ok("编辑友链成功")
		return
	}
}

// Del 删除友链
func (c *LinkController) Del() {
	// 接收参数
	id := c.GetString("id")

	// 业务处理
	ml := models.NewLink()
	dml := ml.GetLinkByUid(id)
	if dml == nil {
		c.Fail("要删除友链不存在")
		return
	}

	err := dml.Delete()
	if err != nil {
		c.Fail("删除友链失败")
		return
	} else {
		c.Fail("删除友链成功")
		return
	}
}
