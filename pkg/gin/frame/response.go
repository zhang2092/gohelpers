package frame

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Response 数据结构体
type response struct {
	// Code 业务状态码
	Code int `json:"code"`
	// Message 提示信息
	Message string `json:"message"`
	// Data 数据，用interface{}的目的是可以用任意数据
	Data interface{} `json:"data"`
	// RequestId 请求ID
	RequestId string `json:"request_id"`
	// Errors 错误提示，如 xx字段不能为空等
	// Errors []ErrorItem `json:"errors"`
}

// ErrorItem 错误项
//type ErrorItem struct {
//	Key   string `json:"key"`
//	Value string `json:"error"`
//}

// PageData 分页数据
type PageData struct {
	Total    int         `json:"total"`
	PageID   int         `json:"page_id"`
	PageSize int         `json:"page_size"`
	Result   interface{} `json:"result"`
}

// NewResponse return response instance
func NewResponse() *response {
	return &response{
		Code:      200,
		Message:   "",
		Data:      nil,
		RequestId: uuid.NewString(),
	}
}

// Wrapper include context
type wrapper struct {
	ctx *gin.Context
}

// WrapContext wrap content
func WrapContext(ctx *gin.Context) *wrapper {
	return &wrapper{ctx: ctx}
}

// Json 输出json,支持自定义response结构体
func (wrapper *wrapper) Json(response *response) {
	wrapper.ctx.JSON(200, response)
}

// Success 成功的输出
func (wrapper *wrapper) Success(message string, data interface{}) {
	response := NewResponse()
	response.Message = message
	response.Data = data
	wrapper.Json(response)
}

// Error 错误输出
func (wrapper *wrapper) Error(code int, message string) {
	response := NewResponse()
	response.Code = code
	response.Message = message
	wrapper.Json(response)
}
