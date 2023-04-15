package admin

// LinkCreateRequest 链接首页入口参数
type LinkCreateRequest struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Sort     int    `json:"sort"`
	Show     int    `json:"show"`
	Describe string `json:"describe"`
}

// LinkCreateResponse 链接首页出口参数
type LinkCreateResponse struct{}

// LinkEditRequest 链接首页入口参数
type LinkEditRequest struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Sort     int    `json:"sort"`
	Show     int    `json:"show"`
	Describe string `json:"describe"`
}

// LinkEditResponse 链接首页出口参数
type LinkEditResponse struct{}
