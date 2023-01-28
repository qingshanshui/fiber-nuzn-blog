package middleware

import (
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
	"net/url"
)

func AdminAuth(ctx *context2.Context) {
	// 取访问的url
	pathname := ctx.Request.URL.String()
	urlPath, _ := url.Parse(pathname) //urlPath.Path  /role/edit
	// 判断 访问的url 是不是 "/admin/login"
	if string(urlPath.Path) != "/admin/login" {
		// 获取token 入股token 解密失败就 清除token跳转到登录页
		cookie := ctx.Input.Cookie("token")
		Secret, _ := beego.AppConfig.String("Secret")
		_, err := utils.ParseToken(cookie, Secret)
		if err != nil {
			ctx.SetCookie("token", "", -1)
			ctx.Redirect(302, "/admin/login")
		}
	}else {
		// 获取token 入股token 解密失败就 清除token跳转到登录页
		cookie := ctx.Input.Cookie("token")
		if cookie != "" {
			Secret, _ := beego.AppConfig.String("Secret")
			_, err := utils.ParseToken(cookie, Secret)
			if err != nil {
				ctx.SetCookie("token", "", -1)
				ctx.Redirect(302, "/admin/login")
			}else {
				ctx.Redirect(302, "/admin/home")
			}
		}
	}
}
