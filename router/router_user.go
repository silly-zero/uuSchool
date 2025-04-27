package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silly-zero/uuSchool/midlleware"
	"github.com/silly-zero/uuSchool/service"
)

func UserRouter(r *gin.Engine) {
	//小程序端-未登录
	mini := r.Group("/mini/user")
	{
		mini.POST("/wxLogin", service.WxLogin)
	}
	//小程序端-已登录
	miniLogin := r.Group("/mini/user")
	miniLogin.Use(midlleware.MidiLoginAuth)
	{
		mini.GET("/testLogin")
	}
	//管理端
	admin := r.Group("/admin/user")
	{
		admin.GET("/test")
	}
}
