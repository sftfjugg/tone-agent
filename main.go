package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"tone-agent/entity"
	"tone-agent/router"
	"tone-agent/service"
)

func main() {
	//创建路由
	ta := gin.Default()
	ta.LoadHTMLGlob("templates/*")
	ta.Static("static", "./static")

	router.AdminRouter(ta)
	router.APIRouter(ta)

	// 监听配置
	viper.AddConfigPath(entity.ConfigDir)
	viper.SetConfigName(entity.APPName)
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Configuration files have been changed, file:", e.Name)
	})

	// 启动定时任务
	service.InitCron()
	// 监听端口
	ta.Run(":8479")
}
