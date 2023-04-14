package NavbarController

import (
	"errors"
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

type NavbarController struct {
	controllers.Base
}

func NewNavbarController() *NavbarController {
	return &NavbarController{}
}

// GetAll 获取全部导航栏
func (t *NavbarController) GetAll(c *fiber.Ctx) error {

	mn := models.NewNavbar()
	navBarAll := mn.GetWebNavBarListAll()
	return c.Render("admin/navBar/index", fiber.Map{
		"navBarAll": navBarAll,
	}, "admin/layout/index")
}

// Add 渲染添加分类
func (t *NavbarController) Add(c *fiber.Ctx) error {
	return c.Render("admin/navBar/add_navBar", fiber.Map{}, "admin/layout/index")
}

// AddPost 添加分类的post
func (t *NavbarController) AddPost(c *fiber.Ctx) error {
	// 接收数据
	title := c.Query("title")
	url := c.Query("url")
	//sort := c.Query("sort")
	//show := c.Query("show")

	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()

	mn := models.NewNavbar()
	mn.Uid = uid
	mn.Title = title
	mn.Url = url
	//mn.Show = show
	//mn.Sort = sort

	// 业务处理
	err := mn.Create()
	if err != nil {
		return c.JSON(t.Fail(errors.New("创建导航栏失败")))

	} else {
		return c.JSON(t.Fail(errors.New("创建导航栏成功")))
	}
}

// Edit 渲染编辑分类
func (t *NavbarController) Edit(c *fiber.Ctx) error {
	// 接收参数
	id := c.Params("id")
	// 业务处理
	mn := models.NewNavbar()
	sort := mn.GetLinkByUid(id)
	return c.Render("admin/navBar/edit_navBar", fiber.Map{
		"Result": sort,
	}, "admin/layout/index")
}

// EditPost 编辑分类的post
func (t *NavbarController) EditPost(c *fiber.Ctx) error {
	// 接收参数
	id := c.Query("id")
	title := c.Query("title")
	url := c.Query("url")
	//sort, _ := c.Query("sort")
	//show, _ := c.Query("show")
	// 组装参数
	mn := models.NewNavbar()
	mn.Title = title
	mn.Url = url
	//mn.Sort = sort
	//mn.Show = show
	// 业务处理
	err := mn.Update(id)
	if err != nil {
		return c.JSON(t.Fail(errors.New("编辑导航栏失败")))
	} else {
		return c.JSON(t.Fail(errors.New("编辑导航栏成功")))
	}
}

// Del 删除分类
func (t *NavbarController) Del(c *fiber.Ctx) error {
	id := c.Params("id")

	mn := models.NewNavbar()
	dmn := mn.GetLinkByUid(id)
	if dmn == nil {
		//c.Fail("删除导航栏不存在")
		//return
		return c.JSON(t.Fail(errors.New("删除导航栏不存在")))
	}
	err := dmn.Delete()
	if err != nil {
		return c.JSON(t.Fail(errors.New("删除导航栏失败")))
	} else {
		return c.JSON(t.Fail(errors.New("删除导航栏成功")))
	}
}
