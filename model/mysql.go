package model

import (
	"github.com/silly-zero/uuSchool/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// DB 数据库初始化
var DB *gorm.DB

func InitMysql() {
	//自定义log日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql的阀值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //颜色
		})
	dsn := config.GConfig.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatal("数据库连接失败")
	}
	log.Println("数据库连接成功")
	DB = db
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	//创建数据库
	//db.AutoMigrate(&UserInfo{})
}
