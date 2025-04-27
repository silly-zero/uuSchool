package response

import "github.com/gin-gonic/gin"

const (
	SUCCESS    = 200
	SuccessMsg = "success"
	ERROR      = 500
	ErrorMsg   = "error"
)

type Res struct {
	Code int    `json:"code"`
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg"`
}

func Response(context *gin.Context, httpStatus int, code int, data any, msg string) {
	context.JSON(httpStatus, Res{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(c *gin.Context, data any) {
	Response(c, SUCCESS, SUCCESS, data, SuccessMsg)
}

func Fail(c *gin.Context, code int, msg string) {
	Response(c, ERROR, code, nil, msg)
}

func BindErr(c *gin.Context) {
	Response(c, ERROR, 500, nil, "绑定参数失败")
}
