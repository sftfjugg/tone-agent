package schedule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/spf13/viper"

	"tone-agent/core"
	"tone-agent/entity"
)

func pullTaskSchedule() error {
	pullTaskUrl := core.GetProxyAPIUrl(entity.AgentAPIPullTask)
	tsn := viper.GetString("tsn")
	sign := core.GetSign()
	data := map[string]string{"tsn": tsn, "sign": sign}
	jsonData, _ := json.Marshal(data)
	client := core.GetHttpClient()
	resp, err := client.Post(pullTaskUrl, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Printf("[pullTaskSchedule] pull task error! error: %s", err)
		return err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[pullTaskSchedule]pull task failed!, error: %s", err.Error())
			return err
		}
		taskResult := &entity.TaskResult{}
		err2 := json.Unmarshal(result, taskResult)
		if err2 != nil {
			log.Printf("[pullTaskSchedule]pull task error!, error: %s", err2.Error())
			return err2
		}
		if len(taskResult.Tasks) > 0 {
			for _, task := range taskResult.Tasks {
				log.Printf("[pullTaskSchedule]task(tid: %s) get ready to running...", task.Tid)
				go core.TaskProcessorByActiveMode(task)
				log.Printf("[pullTaskSchedule]task(tid: %s) running...", task.Tid)
			}
		}
	} else {
		log.Printf("[pullTaskSchedule]pull task failed! status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

func syncResultSchedule() error {
	waitingSyncDir := beego.AppConfig.String("WaitingSyncResultDir")
	files, _ := ioutil.ReadDir(waitingSyncDir)
	for _, f := range files {
		filename := f.Name()
		result, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", waitingSyncDir, filename))
		resultMap := make(map[string]string)
		_ = json.Unmarshal(result, &resultMap)
		var updateData map[string]string
		if resultMap["status"] == entity.TaskCompletedStatus {
			updateData = map[string]string{
				"tid":         resultMap["tid"],
				"status":      resultMap["status"],
				"result":      resultMap["result"],
				"task_pid":    resultMap["taskPid"],
				"error_code":  resultMap["errorCode"],
				"error_msg":   resultMap["errorMsg"],
				"exit_code":   resultMap["exitCode"],
				"finish_time": resultMap["finishTime"],
			}
			go core.SyncResultToProxy(updateData, false, true)
		}
	}
	return nil
}

func revisedData() error {
	waitingSyncDir := beego.AppConfig.String("WaitingSyncResultDir")
	files, _ := ioutil.ReadDir(waitingSyncDir)
	for _, f := range files {
		filename := f.Name()
		result, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", waitingSyncDir, filename))
		resultMap := make(map[string]string)
		_ = json.Unmarshal(result, &resultMap)
		if resultMap["status"] == entity.TaskRunningStatus {
			taskPid, _ := strconv.Atoi(resultMap["taskPid"])
			pidExist := core.CheckPid(taskPid)
			if !pidExist {
				script := resultMap["script"]
				resultMap["status"] = entity.TaskCompletedStatus
				if strings.Contains(script, "reboot") {
					log.Printf("[revisedData] pid no longer exists.revised reboot task status:completed,"+
						" tid:%s", resultMap["pid"])
					resultMap["result"] = "reboot success"
				} else {
					// pid不存在且执行时间在10个小时前的，置为完成（失败）
					beforeTenHourTime := time.Now().Add(-time.Hour * 10).Format("2006-01-02 15:04:05")
					if resultMap["startTime"] < beforeTenHourTime {
						resultMap["errorCode"] = entity.PidNotExistErrorCode
						resultMap["errorMsg"] = entity.PidNotExistErrorMsg
						log.Printf("[revisedData] pid no longer exists.revised task status:completed | tid:%s",
							resultMap["pid"])
					}
				}
				core.WriteResult(resultMap, false)
			}
			log.Printf("[revisedData]revised data schedule. tid:%s | status:%s",
				resultMap["tid"], resultMap["status"])
		}
	}
	return nil
}

func heartbeatSchedule() error{
	err := core.SyncHeartbeatToProxy()
	log.Printf("[HeartbeatSchedule]sync hearbeat request result: %s", err)
	return nil
}

func InitTask() {
	rd := toolbox.NewTask("revisedData", "0/60 * * * * *", revisedData)
	if viper.Get("mode") == "active" {
		pt := toolbox.NewTask("pullTask", "0/1 * * * * *", pullTaskSchedule)
		sr := toolbox.NewTask("syncResult", "0/1 * * * * *", syncResultSchedule)
		hb := toolbox.NewTask("heartbeat", "0/60 * * * * *", heartbeatSchedule)
		toolbox.AddTask("pullTask", pt)
		toolbox.AddTask("syncResult", sr)
		toolbox.AddTask("heartbeat", hb)
	}
	toolbox.AddTask("revisedData", rd)
}
