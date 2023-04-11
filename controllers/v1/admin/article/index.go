package ArticleController

import (
	"bytes"
	"errors"
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jaevor/go-nanoid"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// Add 渲染创建文章页面
func (t *DefaultController) Add(c *fiber.Ctx) error {
	mn := models.NewNavBar()
	sort := mn.GetWebNavBarList()
	if sort == nil {
		return c.JSON(t.Fail(errors.New("分类获取失败")))
	}
	return c.Render("admin/article/add_article", fiber.Map{
		"Sort": sort,
	}, "admin/layout/index")
}

// AddPost 创建文章
func (t *DefaultController) AddPost(c *fiber.Ctx) error {

	// 接收参数
	sort, _ := c.ParamsInt("sort")
	title := c.Params("title")
	content := c.Params("content")
	contentHtml := c.Params("contentHtml")
	pic := c.Params("pic")
	show, _ := c.ParamsInt("show")
	navBarId := c.Params("navBarId")

	// token 解析
	cookie := c.Cookies("token")
	MapClaims, err := utils.ParseToken(cookie, viper.GetString("Jwt.Secret"))
	if err != nil {
		return c.JSON(t.Fail(errors.New("token解析失败")))
	}

	ma := models.NewArticle()
	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	ma.Title = title
	ma.Content = content
	ma.ContentHtml = contentHtml
	ma.NavBarId = navBarId
	ma.Uid = uid
	ma.Pic = pic
	ma.Show = show
	ma.Sort = sort
	ma.UserId = MapClaims["user"].(map[string]interface{})["uid"].(string)
	ma.UserName = MapClaims["user"].(map[string]interface{})["username"].(string)

	// 业务处理
	err = ma.Create()
	if err != nil {
		return c.JSON(t.Fail(errors.New("创建文章失败")))
	} else {
		return c.JSON(t.Fail(errors.New("创建文章成功")))
	}
}

// GetAll 获取文章列表
func (t *DefaultController) GetAll(c *fiber.Ctx) error {
	pageSize, _ := c.ParamsInt("pageSize")     // 一页多少条
	pageNumber, _ := c.ParamsInt("pageNumber") //当前页

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNumber == 0 {
		pageNumber = 1
	}

	ma := models.NewArticle()
	article := ma.GetArticleListAll(pageNumber, pageSize)
	if article == nil {
		return c.JSON(t.Fail(errors.New("全部文章列表获取失败")))
	}
	return c.Render("admin/article/article", fiber.Map{
		"Article": article,
	}, "admin/layout/index")
}

// Edit 渲染修改文章页面
func (t *DefaultController) Edit(c *fiber.Ctx) error {
	// 接收参数
	id := c.Params("id")

	// 文章详情
	ma := models.NewArticle()
	article := ma.GetArticleByUid(id)
	// 分类列表
	mn := models.NewNavBar()
	sort := mn.GetWebNavBarList()

	return c.Render("admin/article/edit_article", fiber.Map{
		"Sort":   sort,
		"Result": article,
	}, "admin/layout/index")
}

// EditPost 修改文章
func (t *DefaultController) EditPost(c *fiber.Ctx) error {
	// 接收参数
	sort, _ := c.ParamsInt("sort")
	title := c.Params("title")
	content := c.Params("content")
	contentHtml := c.Params("contentHtml")
	pic := c.Params("pic")
	show, _ := c.ParamsInt("show")
	navBarId := c.Params("navBarId")
	id := c.Params("id")

	// 解析token
	cookie := c.Cookies("token")
	MapClaims, err := utils.ParseToken(cookie, viper.GetString("Jwt.Secret"))
	if err != nil {
		return c.JSON(t.Fail(errors.New("token解析失败")))
	}

	ma := models.NewArticle()
	// 组装数据
	ma.Title = title
	ma.Content = content
	ma.ContentHtml = contentHtml
	ma.NavBarId = navBarId
	ma.Pic = pic
	ma.Show = show
	ma.Sort = sort
	ma.UserId = MapClaims["user"].(map[string]interface{})["uid"].(string)
	ma.UserName = MapClaims["user"].(map[string]interface{})["username"].(string)

	err = ma.Update(id)
	if err != nil {
		return c.JSON(t.Fail(errors.New("文章修改失败")))
	} else {
		return c.JSON(t.Fail(errors.New("文章修改成功")))
	}
}

// Del 删除文章
func (t *DefaultController) Del(c *fiber.Ctx) error {
	uid := c.Params("id")
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
func (t *DefaultController) Baidu(c *fiber.Ctx) error {
	uid := c.Params("id")
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
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
