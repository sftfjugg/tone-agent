package main

import (
	"io"
	"log"
	"os"
	"path"

	"tone-agent/core"
	_ "tone-agent/routers"
	"tone-agent/schedule"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	toneAgentPath, err := core.GetToneAgentPath()
	if err != nil {
		return
	}
	log.Printf("ToneAgent config path: %v", toneAgentPath)
	viper.AddConfigPath(toneAgentPath)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Read config file failed: %s", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Configuration files have been changed, file:", e.Name)
	})

	resultDirs := make(map[string]string, 4)
	resultDirs["result"] = path.Join(toneAgentPath, viper.GetString("result.ResultFileDir"))
	resultDirs["waiting_sync_result"] = path.Join(toneAgentPath, viper.GetString("result.WaitingSyncResultDir"))
	resultDirs["scripts"] = path.Join(toneAgentPath, viper.GetString("result.TmpScriptFileDir"))
	resultDirs["logs"] = path.Join(toneAgentPath, viper.GetString("result.LogFileDir"))

	setBeegoConfig(resultDirs)

	for key, value := range resultDirs {
		if core.CheckFileIsExist(value) {
			continue
		}
		err = os.MkdirAll(value, 0777)
		if err != nil {
			log.Fatalf("Create %v in %v failed: %v", key, value, err)
			return
		}
	}

	logFilePath := path.Join(resultDirs["logs"], viper.GetString("result.LogFileName"))

	// schedule
	schedule.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()

	// log
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Open log file failed! error: %s", err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetPrefix("[tone-agent]")
	log.SetFlags(log.Ldate | log.Ltime)

	beego.Run()
}

func setBeegoConfig(configs map[string]string) {
	beego.BConfig.AppName = viper.GetString("beego.AppName")
	beego.BConfig.RunMode = viper.GetString("beego.RunMode")
	beego.BConfig.Listen.HTTPAddr = viper.GetString("beego.HttpAddr")
	beego.BConfig.Listen.HTTPPort = viper.GetInt("beego.HttpPort")
	beego.BConfig.WebConfig.StaticDir["/down1"] = viper.GetString("beego.StaticDir")
	beego.BConfig.WebConfig.DirectoryIndex = viper.GetBool("beego.DirectoryIndex")
	beego.BConfig.CopyRequestBody = viper.GetBool("beego.CopyRequestBody")
	_ = beego.AppConfig.Set("TmpScriptFileDir", configs["scripts"])
	_ = beego.AppConfig.Set("WaitingSyncResultDir", configs["waiting_sync_result"])
	_ = beego.AppConfig.Set("ResultFileDir", configs["result"])
}
