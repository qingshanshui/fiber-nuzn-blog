package admin

import "fiber-nuzn-blog/validator/form"

// ArticleHomeRequest 链接首页入口参数
type ArticleHomeRequest struct {
	form.PaginationRequest
}

// ArticleCreateRequest 链接首页入口参数
type ArticleCreateRequest struct {
	Title       string `json:"title"`
	Sort        int    `json:"sort"`
	Show        int    `json:"show"`
	Content     string `json:"content"`
	ContentHtml string `json:"contentHtml"`
	Pic         string `json:"pic"`
	NavbarId    string `json:"navbarId"`
}

// ArticleCreateResponse 链接首页出口参数
type ArticleCreateResponse struct{}

// ArticleEditRequest 链接首页入口参数
type ArticleEditRequest struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Sort        int    `json:"sort"`
	Show        int    `json:"show"`
	Content     string `json:"content"`
	ContentHtml string `json:"contentHtml"`
	Pic         string `json:"pic"`
	NavbarId    string `json:"navbarId"`
}

// ArticleEditResponse 链接首页出口参数
type ArticleEditResponse struct{}
