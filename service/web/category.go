package web

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form/web"
)

type Category struct{}

func NewCategoryService() *Category {
	return &Category{}
}

// Category 详情
func (t *Category) Category(uid string) web.CategoryResponse {
	ma := models.NewArticle()
	// 通过uid 获取详情
	cy := ma.GetArticleByUid(uid)
	if cy != nil {
		// 文章点击数增加
		ma.AddHot()
	}
	return web.CategoryResponse{
		ArticleCategory:     *cy,
		ArticleCategoryTime: cy.UpdatedAt.Format("2006-01-02"),
	}
}
