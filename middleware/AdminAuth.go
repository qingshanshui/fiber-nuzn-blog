package middleware

import (
	"fiber-nuzn-blog/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func AdminAuth(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("token")
	if cookie == "" {
		return ctx.Redirect("/admin/login")
	} else {
		user, err := utils.ParseToken(cookie, viper.GetString("Jwt.Secret"))
		if err != nil {
			return ctx.Redirect("/admin/login")
		} else {
			ctx.Locals("userinfo", user["user"])
			return ctx.Next()
		}
	}
}
