package midlleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(c *gin.Context) {
	method := c.Request.Method
	//允许所有来源的请求进行跨域访问
	c.Header("Access-Control-Allow-Origin", "*")
	//允许的请求头
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,x-token")
	//允许的请求方法
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PATCH,PUT")
	//允许的响应头
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//允许携带凭证
	c.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
}
