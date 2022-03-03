package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func SyncStatusToProxy(tid string, status string) bool {
	syncUrl := GetProxyAPIUrl("SyncResultApi")
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
	syncUrl := GetProxyAPIUrl("SyncResultApi")
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
	syncUrl := fmt.Sprintf("%s", GetProxyAPIUrl("SyncResultApi"))
	tid := values["tid"]
	jsonValue, _ := json.Marshal(values)
	client := GetHttpClient()
	resp, err := client.Post(syncUrl, "application/json", bytes.NewBuffer(jsonValue))
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
				"sync result to proxy faild, tid:%s | detail:%s | status_code:%d",
			tid,
			values,
			resp.StatusCode,
		)
	}
	defer resp.Body.Close()
	return false
}
