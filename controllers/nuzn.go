package controllers

import beego "github.com/beego/beego/v2/server/web"

type Nuzn struct {
	beego.Controller
}

// 返回的 code 码值
const (
	ERROR   = 0     // 失败
	SUCCESS = 20000 // 成功
)

// Ok 成功
func (n *Nuzn) Ok(data interface{}) {
	n.Data["json"] = map[string]interface{}{
		"code": SUCCESS,
		"data": data,
		"msg":  "操作成功",
	}
	n.ServeJSON()
}

// Fail 失败
func (n *Nuzn) Fail(data interface{}) {
	n.Data["json"] = map[string]interface{}{
		"code": ERROR,
		"data": data,
		"msg":  "操作失败",
	}
	n.ServeJSON()
}
