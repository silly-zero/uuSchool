package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//设置项目运行模式
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}
