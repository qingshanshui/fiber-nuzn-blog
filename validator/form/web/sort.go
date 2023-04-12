package web

import "fiber-nuzn-blog/validator/form"

// SortRequest 分类页入口参数
type SortRequest struct {
	form.PaginationRequest
}

// SortResponse 分类页入口参数
type SortResponse struct {
	form.PaginationResponse
}
