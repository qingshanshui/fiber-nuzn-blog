package admin

// LoginRequest 链接首页入口参数
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 链接首页出口参数
type LoginResponse struct{}
