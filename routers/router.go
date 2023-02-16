package routers

import (
	"beego_blog_mvc/controllers/admin"
	ArticleController "beego_blog_mvc/controllers/admin/article"
	LinkController "beego_blog_mvc/controllers/admin/link"
	LoginController "beego_blog_mvc/controllers/admin/login"
	adminSort "beego_blog_mvc/controllers/admin/navBar"
	"beego_blog_mvc/controllers/web"
	CategoryController "beego_blog_mvc/controllers/web/category"
	SortController "beego_blog_mvc/controllers/web/sort"
	"beego_blog_mvc/middleware"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 博客前台
	beego.Router("/", &web.MainController{})
	// 分类详情
	beego.Router("/sort/:id", &SortController.SortController{})
	// 文章详情
	beego.Router("/category/:id([A-Za-z0-9_-]+).html", &CategoryController.CategoryController{}, "get:Category")

	// 博客后台api
	var nsAdmin = beego.NewNamespace("/admin",
		//中间件:匹配路由前会执,可以用于权限验证
		beego.NSBefore(middleware.AdminAuth),
		// 登录login
		beego.NSNamespace("/login",
			beego.NSRouter("/", &LoginController.LoginController{}),
			beego.NSRouter("/logout", &LoginController.LoginController{}, "get:Logout"),
		),

		// 首页
		beego.NSNamespace("/home",
			beego.NSRouter("/", &admin.HomeController{}),
		),

		// 文章管理
		beego.NSNamespace("/article",
			beego.NSRouter("/", &ArticleController.ArticleController{}, "get:GetAll"),
			beego.NSRouter("/add", &ArticleController.ArticleController{}, "get:Add"),
			beego.NSRouter("/add", &ArticleController.ArticleController{}, "post:AddPost"),
			beego.NSRouter("/edit", &ArticleController.ArticleController{}, "get:Edit"),
			beego.NSRouter("/edit", &ArticleController.ArticleController{}, "post:EditPost"),
			beego.NSRouter("/del", &ArticleController.ArticleController{}, "post:Del"),
			beego.NSRouter("/baidu", &ArticleController.ArticleController{}, "post:Baidu"),
		),

		// 导航栏
		beego.NSNamespace("/navbar",
			beego.NSRouter("/", &adminSort.SortController{}, "get:GetAll"),
			beego.NSRouter("/add", &adminSort.SortController{}, "get:Add"),
			beego.NSRouter("/add", &adminSort.SortController{}, "post:AddPost"),
			beego.NSRouter("/edit", &adminSort.SortController{}, "get:Edit"),
			beego.NSRouter("/edit", &adminSort.SortController{}, "post:EditPost"),
			beego.NSRouter("/del", &adminSort.SortController{}, "post:Del"),
		),

		// 友链
		beego.NSNamespace("/link",
			beego.NSRouter("/", &LinkController.LinkController{}, "get:GetAll"),
			beego.NSRouter("/add", &LinkController.LinkController{}, "get:Add"),
			beego.NSRouter("/add", &LinkController.LinkController{}, "post:AddPost"),
			beego.NSRouter("/edit", &LinkController.LinkController{}, "get:Edit"),
			beego.NSRouter("/edit", &LinkController.LinkController{}, "post:EditPost"),
			beego.NSRouter("/del", &LinkController.LinkController{}, "post:Del"),
		),
	)
	beego.AddNamespace(nsAdmin)
}
