package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"tone-agent/core"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"tone-agent/controllers"
	"tone-agent/schedule"
)


func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("toneagent.config.yaml")
	viper.AddConfigPath("/usr/local/toneagent/conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Read config file failed: %s", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Configuration files have been changed, file:", e.Name)
	})

	// make dir
	resultFileDir := beego.AppConfig.String("ResultFileDir")
	waitingSyncDir := beego.AppConfig.String("WaitingSyncResultDir")
	if !core.CheckFileIsExist(resultFileDir){
		core.MakeDir(resultFileDir)
	}
	if !core.CheckFileIsExist(waitingSyncDir){
		core.MakeDir(waitingSyncDir)
	}
	tmpScriptDir := beego.AppConfig.String("TmpScriptFileDir")
	if !core.CheckFileIsExist(tmpScriptDir){
		core.MakeDir(tmpScriptDir)
	}

	logFileDir := beego.AppConfig.String("LogFileDir")
	if !core.CheckFileIsExist(logFileDir){
		core.MakeDir(logFileDir)
	}
	logFileName := beego.AppConfig.String("LogFileName")
	logFilePath := fmt.Sprintf("%s/%s", logFileDir, logFileName)

	// schedule
	schedule.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()

	// log
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Open log file failed! error: %s", err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout,logFile)
	log.SetOutput(mw)
	log.SetPrefix("[tone-agent]")
	log.SetFlags(log.Ldate|log.Ltime)
	//beego.BConfig.CopyRequestBody = true
	// router
	beego.Router("api/task", &controllers.TaskController{})
	beego.Router("api/query", &controllers.ResultController{})
	beego.Router("api/heartbeat", &controllers.HeartbeatController{})

	beego.Run()
}