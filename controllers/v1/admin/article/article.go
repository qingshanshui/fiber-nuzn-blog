package ArticleController

import (
	"bytes"
	"errors"
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	serviceAdmin "fiber-nuzn-blog/service/admin"
	"fiber-nuzn-blog/validator"
	validatorForm "fiber-nuzn-blog/validator/form/admin"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"time"
)

type ArticleController struct {
	controllers.Base
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// Home 获取文章列表
func (t *ArticleController) Home(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleHomeRequestForm := validatorForm.ArticleHomeRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &ArticleHomeRequestForm); err != nil {
		return err
	}
	// 分页调用
	t.PaginationInit(&ArticleHomeRequestForm.PaginationRequest)
	r := serviceAdmin.NewArticleService().Home(ArticleHomeRequestForm.PaginationRequest)
	return c.Render("admin/article/article", fiber.Map{
		"Article": r,
	}, "admin/layout/index")
}

// AddView 渲染创建文章页面
func (t *ArticleController) AddView(c *fiber.Ctx) error {
	r := serviceAdmin.NewArticleService().AddView()
	return c.Render("admin/article/create", fiber.Map{
		"Sort": r,
	}, "admin/layout/index")
}

// Add 创建文章
func (t *ArticleController) Add(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleCreateRequestForm := validatorForm.ArticleCreateRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleCreateRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	err := serviceAdmin.NewArticleService().Add(ArticleCreateRequestForm, c)
	if err != nil {
		return c.JSON(t.Fail(errors.New("创建文章失败")))
	} else {
		return c.JSON(t.Fail(errors.New("创建文章成功")))
	}
}

// EditView 渲染修改文章页面
func (t *ArticleController) EditView(c *fiber.Ctx) error {
	// 接收参数
	id := c.FormValue("id")
	// 实际业务调用
	r := serviceAdmin.NewArticleService().EditView(id)
	return c.Render("admin/article/edit", fiber.Map{
		"Sort":   r.Navbar,
		"Result": r.Article,
	}, "admin/layout/index")
}

// Edit 修改文章
func (t *ArticleController) Edit(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleEditRequestForm := validatorForm.ArticleEditRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleEditRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	err := serviceAdmin.NewArticleService().Edit(ArticleEditRequestForm, c)
	if err != nil {
		return c.JSON(t.Fail(errors.New("文章修改失败")))
	} else {
		return c.JSON(t.Fail(errors.New("文章修改成功")))
	}
}

// Del 删除文章
func (t *ArticleController) Del(c *fiber.Ctx) error {
	uid := c.FormValue("id")
	ma := models.NewArticle()
	dma := ma.GetArticleByUid(uid)
	err := dma.Delete()
	if err != nil {
		return c.JSON(t.Fail(errors.New("删除文章失败")))
	} else {
		return c.JSON(t.Fail(errors.New("删除文章成功")))
	}
}

// Baidu 百度推送
func (t *ArticleController) Baidu(c *fiber.Ctx) error {
	uid := c.FormValue("id")
	content := Post("http://data.zz.baidu.com/urls?site=blog.nuzn.cn&token=iW9On1j50GCOJUxq", "https://blog.nuzn.cn/category/"+uid+".html", "text/plain")
	return c.JSON(content)
}

func Post(url string, data interface{}, contentType string) (content string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.(string))))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, errorClient := client.Do(req)
	if errorClient != nil {
		panic(errorClient)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
