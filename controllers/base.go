package controllers

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
	"github.com/gofiber/fiber/v2"
)

// 返回的 code 码值
const (
	ERROR   = 0
	SUCCESS = 1000
)

// Base 公共方法
type Base struct{}

// Ok 成功
func (r *Base) Ok(data interface{}) fiber.Map {
	return fiber.Map{
		"code": SUCCESS,
		"data": data,
		"msg":  "操作成功",
	}
}

// Fail 失败
func (r *Base) Fail(err error, code ...int) fiber.Map {
	return fiber.Map{
		"code": If(code),
		"data": err.Error(),
		"msg":  "操作失败",
	}
}

func If(code []int) int {
	if code != nil {
		return code[0]
	}
	return ERROR
}

// InitData 首页查询接口
func (r *Base) InitData() form.InitData {
	// 导航栏
	mn := models.NewNavBar()
	nl := mn.GetWebNavBarList()

	// 设置cookie
	//cookie := c.Ctx.Input.Cookie("token")

	/*        站点统计      */
	ma := models.NewArticle()
	// 获取文章总条数
	articleTotalCount := ma.GetArticleCount()

	// 分类数（navBar表）
	navBarCount := mn.GetNavBarListCount()

	// 页面数（article表+navBar表）
	pagesCount := articleTotalCount + navBarCount

	// 更新 （文章表最后一条内容）
	ut := ma.GetArticleLast()

	// 获取友链（友情链接）
	ml := models.NewLink()
	ll := ml.GetLinkList()

	return form.InitData{
		ArticleTotalCount: articleTotalCount,
		NavBarCount:       navBarCount,
		PagesCount:        pagesCount,
		UpdateTime:        ut.UpdatedAt.Format("2006-01-02"),
		LinkAllList:       ll,
		WebNavBarList:     nl,
	}
	//c.Data["Cookie"] = cookie
}

// PaginationInit 初始化分页
func (r *Base) PaginationInit(request *form.PaginationRequest) {
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	if request.CurrPage == 0 {
		request.CurrPage = 1
	}
}
