package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	_ "tone-agent/routers"
	"tone-agent/schedule"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func initConfig()  {
	viper.AddConfigPath(".")
	viper.AddConfigPath(beego.AppConfig.String("AgentConfigFilePath"))
	viper.SetConfigName(beego.AppConfig.String("AgentConfigFileName"))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Read config file failed: %s", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Configuration files have been changed, file:", e.Name)
	})
}

func initLog()  {
	// log
	logFileDir := beego.AppConfig.String("LogFileDir")
	logFileName := beego.AppConfig.String("LogFileName")
	logFilePath := fmt.Sprintf("%s/%s", logFileDir, logFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Open log file failed! error: %s", err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetPrefix("[tone-agent]")
	log.SetFlags(log.Ldate | log.Ltime)
}


func main() {
	initConfig()

	schedule.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()

	initLog()
	beego.Run()
}