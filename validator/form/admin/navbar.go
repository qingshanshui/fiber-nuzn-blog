package admin

// NavbarRequest 分类首页入口参数
type NavbarRequest struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Sort  int    `json:"sort"`
	Show  int    `json:"show"`
}

// NavbarCreateResponse 分类首页出口参数
type NavbarCreateResponse struct {
}

// NavbarEditRequest 分类首页入口参数
type NavbarEditRequest struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Sort  int    `json:"sort"`
	Show  int    `json:"show"`
}

// NavbarEditResponse 分类首页出口参数
type NavbarEditResponse struct{}
