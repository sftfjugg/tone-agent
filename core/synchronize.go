package core

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
	"tone-agent/entity"
)

func SyncStatusToProxy(tid string, status string) bool {
	syncUrl := GetProxyAPIUrl(entity.AgentAPISyncResult)
	sign := GetSign()
	values := map[string]string{"tid": tid, "status": status, "sign": sign}
	jsonValue, _ := json.Marshal(values)
	client := GetHttpClient()
	resp, err := client.Post(syncUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf(
			"[SyncStatusToProxy]"+
				"sync status to proxy error! tid:%s | status:%s | error: %s",
			tid,
			status,
			err,
		)
		return false
	}
	if resp.StatusCode == 200 {
		log.Printf(
			"[SyncStatusToProxy]"+
				"sync status to proxy success, tid:%s | status:%s",
			tid,
			status,
		)
		defer resp.Body.Close()
		return true
	} else {
		log.Printf(
			"[SyncStatusToProxy]"+
				"sync status to proxy failed, tid:%s | status code:%d",
			tid,
			resp.StatusCode,
		)
		defer resp.Body.Close()
		return false
	}
}

func SyncExecTimeToProxy(tid string, timeType string, pid string) bool {
	syncUrl := GetProxyAPIUrl(entity.AgentAPISyncResult)
	sign := GetSign()
	values := map[string]string{
		"tid":      tid,
		"task_pid": pid,
		timeType:   time.Now().Format("2006-01-02 15:04:05"),
		"sign":     sign,
	}
	jsonValue, _ := json.Marshal(values)
	client := GetHttpClient()
	resp, err := client.Post(syncUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf(
			"[SyncExecTimeToProxy]"+
				"sync exec time to proxy error, tid:%s | timeType:%s | error:%s",
			tid,
			timeType,
			err.Error(),
		)
		return false
	}
	if resp.StatusCode == 200 {
		log.Printf(
			"[SyncExecTimeToProxy]"+
				"sync exec time to proxy success, tid:%s | timeType:%s",
			tid,
			timeType,
		)
		defer resp.Body.Close()
		return true
	} else {
		log.Printf(
			"[SyncStatusToProxy]sync exec time to proxy failed, tid:%s | status code:%d",
			tid,
			resp.StatusCode,
		)
		defer resp.Body.Close()
	}
	return false
}

func SyncResultToProxy(values map[string]string, sync bool, finish bool) bool {
	values["sign"] = GetSign()
	tid := values["tid"]
	jsonValue, _ := json.Marshal(values)
	client := GetHttpClient()
	resp, err := client.Post(GetProxyAPIUrl(entity.AgentAPISyncResult), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf(
			"[SyncResultToProxy]sync result to proxy error, tid:%s | error:%s",
			tid,
			err.Error(),
		)
		return false
	}
	if resp.StatusCode == 200 {
		if !sync && finish {
			filename := GetFileNameByTid(tid)
			ReMoveFile(filename)
		}
		log.Printf(
			"[SyncResultToProxy]sync result to proxy success, tid:%s | status:%s",
			tid,
			values["status"],
		)
		defer resp.Body.Close()
		return true
	} else {
		log.Printf(
			"[SyncResultToProxy]"+
				"sync result to proxy faild, tid:%s | status_code:%d",
			tid,
			resp.StatusCode,
		)
	}
	defer resp.Body.Close()
	return false
}

func SyncHeartbeatToProxy() string {
	heartbeatAPI := GetProxyAPIUrl(entity.AgentAPIHeartbeat)
	tsn := viper.GetString("tsn")
	sign := GetSign()
	arch, _ := ExecCommand("arch")
	kernel, _ := ExecCommand("uname -r")
	distro, _ := ExecCommand("cat /etc/os-release | grep -i id=")
	data := map[string]string{
		"tsn":    tsn,
		"sign":   sign,
		"arch":   arch,
		"kernel": kernel,
		"distro": distro,
	}
	jsonData, _ := json.Marshal(data)
	client := GetHttpClient()
	resp, err := client.Post(heartbeatAPI, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Printf("[HeartbeatSchedule]Hearbeat info sync error, url:%s | error:%s", heartbeatAPI, err.Error())
		return err.Error()
	}
	result, _ := ioutil.ReadAll(resp.Body)
	var resData map[string]interface{}
	err = json.Unmarshal([]byte(result), &resData)
	if err != nil{
		return err.Error()
	}
	if resData["SUCCESS"] == "FALSE" {
		errorMsg := resData["ERROR_MSG"]
		return errorMsg.(string)
	}
	if resp.StatusCode != 200 {
		log.Printf("[HeartbeatSchedule]Hearbeat schedule request failed! StatusCode:%d | detail:%s",
			resp.StatusCode, result)
		return err.Error()
	}
	defer resp.Body.Close()
	return ""
}