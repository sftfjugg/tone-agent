package service

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"strings"
	"tone-agent/entity"
)

func heartbeat() {
	SyncHeartbeat()
}

func pullTask() {
	tasks := PullTask()
	ProcessTask(tasks)
}

func syncResult() {
	files := ScanResult()
	for _, f := range files {
		tid := strings.Trim(f, ".json")
		result, _ := ReadResult(tid)
		SyncResult(result)
	}

}

func InitCron() {
	c := cron.New(cron.WithSeconds())
	if viper.Get("mode") == "active" {
		c.AddFunc(entity.HeartbeatInterval, heartbeat)
		c.AddFunc(entity.PullTaskInterval, pullTask)
		c.AddFunc(entity.SyncResultInterval, syncResult)
	}
	c.Start()
}
