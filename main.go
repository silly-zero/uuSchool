package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/silly-zero/uuSchool/configs"
	"github.com/silly-zero/uuSchool/model"
	"github.com/silly-zero/uuSchool/router"
	"log"
)

func main() {
	//设置项目运行模式
	gin.SetMode(gin.DebugMode)
	//初始化配置文件
	config.InitConfig()
	//初始化路由
	r := router.InitRouter()
	//初始化Mysql
	model.InitMysql()
	address := config.GConfig.System.GetAddress()
	err := r.Run(address)
	if err != nil {
		log.Fatal("项目启动失败！！！")
		return
	}
}
