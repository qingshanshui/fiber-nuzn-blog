package LinkController

import (
	"errors"
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"

	"github.com/jaevor/go-nanoid"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// GetAll 获取友链
func (t *DefaultController) GetAll(c *fiber.Ctx) error {
	ml := models.NewLink()
	linkAll := ml.GetAdminLinkList()
	return c.Render("admin/link/index", fiber.Map{
		"linkAll": linkAll,
	}, "admin/layout/index")
}

// Add 渲染添加友链
func (t *DefaultController) Add(c *fiber.Ctx) error {
	return c.Render("admin/link/add_link", fiber.Map{}, "admin/layout/index")
}

// AddPost 添加友链的post
func (t *DefaultController) AddPost(c *fiber.Ctx) error {
	// 接收参数
	title := c.Params("title")
	url := c.Params("url")
	sort, _ := c.ParamsInt("sort")
	show, _ := c.ParamsInt("show")
	describe := c.Params("describe")

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
		return c.JSON(t.Fail(errors.New("创建友链失败")))
	} else {
		return c.JSON(t.Fail(errors.New("创建友链成功")))
	}
}

// Edit 渲染 编辑友链
func (t *DefaultController) Edit(c *fiber.Ctx) error {

	id := c.Params("id")
	ml := models.NewLink()
	sort := ml.GetLinkByUid(id)
	return c.Render("admin/link/edit_link", fiber.Map{
		"Result": sort,
	}, "admin/layout/index")
}

// EditPost 编辑友链的post
func (t *DefaultController) EditPost(c *fiber.Ctx) error {
	// 接收参数
	id := c.Params("id")
	title := c.Params("title")
	url := c.Params("url")
	sort, _ := c.ParamsInt("sort")
	show, _ := c.ParamsInt("show")
	describe := c.Params("describe")

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
		return c.JSON(t.Fail(errors.New("编辑友链失败")))
	} else {
		return c.JSON(t.Fail(errors.New("编辑友链成功")))
	}
}

// Del 删除友链
func (t *DefaultController) Del(c *fiber.Ctx) error {
	// 接收参数
	id := c.Params("id")

	// 业务处理
	ml := models.NewLink()
	dml := ml.GetLinkByUid(id)
	if dml == nil {
		return c.JSON(t.Fail(errors.New("删除友链不存在")))
	}
	err := dml.Delete()
	if err != nil {
		return c.JSON(t.Fail(errors.New("删除友链失败")))
	} else {
		return c.JSON(t.Fail(errors.New("删除友链成功")))
	}
}
