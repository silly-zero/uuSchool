package config

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Wechat Wechat `yaml:"wechat"`
}

var GConfig *Config

func InitConfig() {
	var configPath string
	mode := gin.Mode()
	if mode == gin.DebugMode {
		configPath = "configs/config_dev.yaml"
	} else if mode == gin.ReleaseMode {
		configPath = "configs/config_prod.yaml"
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("配置文件读取失败！！！", err.Error())
	}
	c := new(Config)
	err = yaml.Unmarshal(file, c)
	GConfig = c
	if err != nil {
		log.Fatal("配置文件解析失败！！！", err.Error())
	}
	log.Println("配置文件初始化成功！！！")
}
