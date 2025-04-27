package midlleware

import (
	"github.com/gin-gonic/gin"
	"github.com/silly-zero/uuSchool/response"
	"github.com/silly-zero/uuSchool/utils"
)

// 登录验证
func MidiLoginAuth(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	claims, err := util.ParseToken(token)
	if err != nil {
		response.Fail(ctx, 401, "认证失败")
		ctx.Abort()
		return
	}
	ctx.Set("userID", claims.ID)
}
