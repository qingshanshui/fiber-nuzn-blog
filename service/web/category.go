package web

import (
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
)

type Category struct{}

func NewCategoryService() *Category {
	return &Category{}
}

// Category 详情
func (t *Category) Category(uid string) fiber.Map {

	ma := models.NewArticle()
	// 通过uid 获取详情
	cy := ma.GetArticleByUid(uid)
	if len(cy) != 0 {
		// 文章点击数增加
		ma.AddHot()
		return fiber.Map{
			"Cy": cy[0],
			"Cm": cy[0].UpdatedAt.Format("2006-01-02"),
		}
	}
	return nil
}
