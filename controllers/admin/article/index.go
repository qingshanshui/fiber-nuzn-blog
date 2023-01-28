package ArticleController

import (
	"beego_blog_mvc/controllers"
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
	"bytes"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/jaevor/go-nanoid"
	"io/ioutil"
	"net/http"
	"time"
)

type ArticleController struct {
	controllers.Nuzn
}

// Add 渲染创建文章页面
func (c *ArticleController) Add() {
	mn := models.NewNavBar()
	sort := mn.GetWebNavBarList()
	if sort == nil {
		c.Fail("分类获取失败")
		return
	}
	c.Data["Sort"] = sort

	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/article/add_article.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/article/css.html"
	c.LayoutSections["LayoutJs"] = "admin/article/js.html"
}

// AddPost 创建文章
func (c *ArticleController) AddPost() {

	// 接收参数
	sort, _ := c.GetInt("sort")
	title := c.GetString("title")
	content := c.GetString("content")
	contentHtml := c.GetString("contentHtml")
	pic := c.GetString("pic")
	show, _ := c.GetInt("show")
	navBarId := c.GetString("navBarId")

	// token 解析
	cookie := c.Ctx.Input.Cookie("token")
	Secret, _ := beego.AppConfig.String("Secret")
	MapClaims, err := utils.ParseToken(cookie, Secret)
	if err != nil {
		c.Fail("token解析失败")
		return
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
		c.Fail("创建文章失败")
		return
	} else {
		c.Ok("创建文章成功")
		return
	}
}

// GetAll 获取文章列表
func (c *ArticleController) GetAll() {

	pageSize, _ := c.GetInt("pageSize")     // 一页多少条
	pageNumber, _ := c.GetInt("pageNumber") //当前页

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNumber == 0 {
		pageNumber = 1
	}

	ma := models.NewArticle()
	article := ma.GetArticleListAll(pageNumber, pageSize)
	if article == nil {
		c.Fail("全部文章列表获取失败")
		return
	}
	c.Data["Article"] = article
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/article/article.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/article/css.html"
	c.LayoutSections["LayoutJs"] = "admin/article/js.html"
}

// Edit 渲染修改文章页面
func (c *ArticleController) Edit() {
	// 接收参数
	id := c.GetString("id")

	// 文章详情
	ma := models.NewArticle()
	article := ma.GetArticleByUid(id)
	// 分类列表
	mn := models.NewNavBar()
	sort := mn.GetWebNavBarList()

	// 返回参数
	c.Data["Sort"] = sort
	c.Data["Result"] = article
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/article/edit_article.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/article/css.html"
	c.LayoutSections["LayoutJs"] = "admin/article/js.html"
}

// EditPost 修改文章
func (c *ArticleController) EditPost() {
	// 接收参数
	sort, _ := c.GetInt("sort")
	title := c.GetString("title")
	content := c.GetString("content")
	contentHtml := c.GetString("contentHtml")
	pic := c.GetString("pic")
	show, _ := c.GetInt("show")
	navBarId := c.GetString("navBarId")
	id := c.GetString("id")

	// 解析token
	cookie := c.Ctx.Input.Cookie("token")
	Secret, _ := beego.AppConfig.String("Secret")
	MapClaims, err := utils.ParseToken(cookie, Secret)
	if err != nil {
		c.Fail("token解析失败")
		return
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
		c.Fail("文章修改失败")
		return
	} else {
		c.Fail("文章修改成功")
		return
	}
}

// Del 删除文章
func (c *ArticleController) Del() {
	uid := c.GetString("id")
	ma := models.NewArticle()
	dma := ma.GetArticleByUid(uid)
	err := dma.Delete()
	if err != nil {
		c.Fail("删除文章失败")
		return
	} else {
		c.Ok("删除文章成功")
		return
	}
}

// Baidu 百度推送
func (c *ArticleController) Baidu() {
	uid := c.GetString("id")
	content := Post("http://data.zz.baidu.com/urls?site=blog.nuzn.cn&token=iW9On1j50GCOJUxq", "https://blog.nuzn.cn/category/"+uid+".html", "text/plain")
	fmt.Println(content, "content", uid)
	c.Ok(content)
	return
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
