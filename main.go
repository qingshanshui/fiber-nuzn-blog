package main

import (
	_ "beego_blog_mvc/routers"
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.AddFuncMap("index", Index)
	beego.AddFuncMap("UnixToDate", UnixToDate)
	beego.Run()
}

// UnixToDate 时间戳转日期
func UnixToDate(unix int64) string {
	return utils.UnixDate(unix)
}

// Index 首页自定义序号
func Index(in int) (out int) {
	out = in + 1
	return out
}
