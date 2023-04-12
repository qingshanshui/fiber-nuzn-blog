package routers

import (
	"fiber-nuzn-blog/controllers/v1/admin"
	ArticleController "fiber-nuzn-blog/controllers/v1/admin/article"
	LinkController "fiber-nuzn-blog/controllers/v1/admin/link"
	LoginController "fiber-nuzn-blog/controllers/v1/admin/login"
	NavBarController "fiber-nuzn-blog/controllers/v1/admin/navBar"
	web2 "fiber-nuzn-blog/controllers/v1/web"
	CategoryController "fiber-nuzn-blog/controllers/v1/web/category"
	SortController "fiber-nuzn-blog/controllers/v1/web/sort"
	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	webRouter := app.Group("/")
	{
		home := web2.NewHomeController()
		webRouter.Get("/", home.Home) // 首页
		sort := SortController.NewSortController()
		webRouter.Get("/sort/:id", sort.Sort) // 分类页面
		category := CategoryController.NewCategoryController()
		webRouter.Get("/category/:id.html", category.Category) // 详情
	}
	// 博客后台api
	adminRouter := app.Group("/admin")
	{
		//中间件:匹配路由前会执,可以用于权限验证
		//adminRouter.Use(middleware.AdminAuth)
		// 登录login
		loginAdminRouter := adminRouter.Group("login")
		{
			login := LoginController.NewDefaultController()
			loginAdminRouter.Get("/", login.Get)          // 登录页面
			loginAdminRouter.Post("/", login.Post)        // 登录
			loginAdminRouter.Get("/logout", login.Logout) // 退出
		}
		// 首页
		homeAdminRouter := adminRouter.Group("home")
		{
			adminHome := admin.NewDefaultController()
			homeAdminRouter.Get("/", adminHome.Home) // 首页
		}
		// 文章
		articleAdminRouter := adminRouter.Group("article")
		{
			articleAdmin := ArticleController.NewDefaultController()
			articleAdminRouter.Get("/", articleAdmin.GetAll)        // 文章首页
			articleAdminRouter.Get("/add", articleAdmin.Add)        // 添加页面
			articleAdminRouter.Post("/add", articleAdmin.AddPost)   // 添加接口
			articleAdminRouter.Get("/edit", articleAdmin.Edit)      // 编辑页面
			articleAdminRouter.Post("/edit", articleAdmin.EditPost) // 编辑接口
			articleAdminRouter.Post("/del", articleAdmin.Del)       // 删除接口
			articleAdminRouter.Post("/baidu", articleAdmin.Baidu)   // 百度推送接口
		}
		//导航栏
		navbarAdminRouter := adminRouter.Group("navbar")
		{
			navbarAdmin := NavBarController.NewDefaultController()
			navbarAdminRouter.Get("/", navbarAdmin.GetAll)        // 导航栏首页
			navbarAdminRouter.Get("/add", navbarAdmin.Add)        // 添加页面
			navbarAdminRouter.Post("/add", navbarAdmin.AddPost)   // 添加接口
			navbarAdminRouter.Get("/edit", navbarAdmin.Edit)      // 编辑页面
			navbarAdminRouter.Post("/edit", navbarAdmin.EditPost) // 编辑接口
			navbarAdminRouter.Post("/del", navbarAdmin.Del)       // 删除接口
		}
		//友链
		linkAdminRouter := adminRouter.Group("link")
		{
			linkAdmin := LinkController.NewDefaultController()
			linkAdminRouter.Get("/", linkAdmin.GetAll)        // 友链首页
			linkAdminRouter.Get("/add", linkAdmin.Add)        // 添加页面
			linkAdminRouter.Post("/add", linkAdmin.AddPost)   // 添加接口
			linkAdminRouter.Get("/edit", linkAdmin.Edit)      // 编辑页面
			linkAdminRouter.Post("/edit", linkAdmin.EditPost) // 编辑接口
			linkAdminRouter.Post("/del", linkAdmin.Del)       // 删除接口
		}
	}
}
