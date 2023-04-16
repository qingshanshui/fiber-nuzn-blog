package admin

import (
	"errors"
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
	validatorForm "fiber-nuzn-blog/validator/form/admin"
	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

type Article struct{}

func NewArticleService() *Article {
	return &Article{}
}

// Home 首页
func (t *Article) Home(r form.PaginationRequest) []models.Article {
	ma := models.NewArticle()
	l := ma.GetArticleListAll(r.CurrPage, r.PageSize)
	return l
}

// AddView 添加文章首页
func (t *Article) AddView() []models.NavBar {
	mn := models.NewNavbar()
	l := mn.GetWebNavBarList()
	return l
}

// Add 添加分类的post
func (t *Article) Add(r validatorForm.ArticleCreateRequest, ctx *fiber.Ctx) error {
	ma := models.NewArticle()
	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	ma.Title = r.Title
	ma.Content = r.Content
	ma.ContentHtml = r.ContentHtml
	ma.NavBarId = r.NavbarId
	ma.Uid = uid
	ma.Pic = r.Pic
	ma.Show = r.Show
	ma.Sort = r.Sort
	userinfo := ctx.Locals("userinfo").(map[string]interface{})
	ma.UserId = userinfo["uid"].(string)
	ma.UserName = userinfo["username"].(string)
	// 业务处理
	err := ma.Create()
	if err != nil {
		return err
	}
	return nil
}

func (t *Article) EditView(id string) validatorForm.ArticleEditView {
	r := validatorForm.ArticleEditView{}
	// 文章详情
	ma := models.NewArticle()
	a := ma.GetArticleByUid(id)
	// 分类列表
	mn := models.NewNavbar()
	s := mn.GetWebNavBarList()
	r.Article = a
	r.Navbar = s
	return r
}

// Edit 编辑分类的post
func (t *Article) Edit(r validatorForm.ArticleEditRequest, ctx *fiber.Ctx) error {
	ma := models.NewArticle()
	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	ma.Title = r.Title
	ma.Content = r.Content
	ma.ContentHtml = r.ContentHtml
	ma.NavBarId = r.NavbarId
	ma.Uid = uid
	ma.Pic = r.Pic
	ma.Show = r.Show
	ma.Sort = r.Sort
	userinfo := ctx.Locals("userinfo").(map[string]interface{})
	ma.UserId = userinfo["uid"].(string)
	ma.UserName = userinfo["username"].(string)
	err := ma.Update(r.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Article) Del(id string) error {
	// 业务处理
	ml := models.NewArticle()
	dml := ml.GetArticleByUid(id)
	if dml == nil {
		return errors.New("删除的友链不存在")
	}
	err := dml.Delete()
	if err != nil {
		return err
	}
	return nil
}
