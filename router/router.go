package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silly-zero/uuSchool/midlleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//跨域
	r.Use(midlleware.Cors)
	UserRouter(r)
	return r
}
