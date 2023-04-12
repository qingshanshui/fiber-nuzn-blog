package web

import "fiber-nuzn-blog/validator/form"

// HomeRequest 主页入口参数
type HomeRequest struct {
	form.PaginationRequest
}

// HomeResponse 主页入口参数
type HomeResponse struct {
	form.PaginationResponse
}
