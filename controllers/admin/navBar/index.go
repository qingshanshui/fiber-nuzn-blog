package SortController

import (
	"beego_blog_mvc/controllers"
	"beego_blog_mvc/models"

	"github.com/jaevor/go-nanoid"
)

type SortController struct {
	controllers.Nuzn
}

// GetAll 获取全部导航栏
func (c *SortController) GetAll() {
	mn := models.NewNavBar()
	navBarAll := mn.GetWebNavBarListAll()
	c.Data["navBarAll"] = navBarAll
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/navBar/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/navBar/css.html"
	c.LayoutSections["LayoutJs"] = "admin/navBar/js.html"
}

// Add 渲染添加分类
func (c *SortController) Add() {
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/navBar/add_navBar.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/navBar/css.html"
	c.LayoutSections["LayoutJs"] = "admin/navBar/js.html"
}

// AddPost 添加分类的post
func (c *SortController) AddPost() {

	// 接收数据
	title := c.GetString("title")
	url := c.GetString("url")
	sort, _ := c.GetInt("sort")
	show, _ := c.GetInt("show")

	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()

	mn := models.NewNavBar()
	mn.Uid = uid
	mn.Title = title
	mn.Url = url
	mn.Show = show
	mn.Sort = sort

	// 业务处理
	err := mn.Create()
	if err != nil {
		c.Fail("创建导航栏失败")
		return
	} else {
		c.Ok("创建导航栏成功")
		return
	}
}

// Edit 渲染 编辑分类
func (c *SortController) Edit() {
	// 接收参数
	id := c.GetString("id")
	// 业务处理
	mn := models.NewNavBar()
	sort := mn.GetLinkByUid(id)
	// 返回参数
	c.Data["Result"] = sort
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/navBar/edit_navBar.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/navBar/css.html"
	c.LayoutSections["LayoutJs"] = "admin/navBar/js.html"
}

// EditPost 编辑分类的post
func (c *SortController) EditPost() {
	// 接收参数
	id := c.GetString("id")
	title := c.GetString("title")
	url := c.GetString("url")
	sort, _ := c.GetInt("sort")
	show, _ := c.GetInt("show")
	// 组装参数
	mn := models.NewNavBar()
	mn.Title = title
	mn.Url = url
	mn.Sort = sort
	mn.Show = show
	// 业务处理
	err := mn.Update(id)
	if err != nil {
		c.Fail("编辑导航栏失败")
		return
	} else {
		c.Ok("编辑导航栏成功")
		return
	}
}

// Del 删除分类
func (c *SortController) Del() {
	id := c.GetString("id")

	mn := models.NewNavBar()
	dmn := mn.GetLinkByUid(id)
	if dmn == nil {
		c.Fail("要删除导航栏不存在")
		return
	}
	err := dmn.Delete()
	if err != nil {
		c.Fail("删除导航栏失败")
		return
	} else {
		c.Fail("删除导航栏成功")
		return
	}
}
