package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"tone-agent/entity"
)

func ScanResult() []string {
	files, _ := ioutil.ReadDir(entity.ResultDir)
	var filesList []string

	for _, f := range files {
		filename := f.Name()
		if strings.HasSuffix(filename, ".json") {
			filesList = append(filesList, filename)
		}
	}
	return filesList
}

func ReadResult(tid string) (map[string]string, bool) {
	file := fmt.Sprintf("%s/%s.json", entity.ResultDir, tid)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, false
	}
	result, _ := ioutil.ReadFile(file)
	resultMap := make(map[string]string)
	err2 := json.Unmarshal(result, &resultMap)
	if err2 != nil {
		return nil, false
	}
	return resultMap, true
}

func WriteResult(resultData map[string]string) {
	file := fmt.Sprintf("%s/%s.json", entity.ResultDir, resultData["tid"])
	content, _ := json.Marshal(resultData)
	err := ioutil.WriteFile(file, content, 0666)
	if err != nil {
		//	TODO
	}
}

func UpdateResult(updateData map[string]string) error {
	//var updateData map[string]string
	updateData = map[string]string{
		"tid":         updateData["tid"],
		"status":      updateData["status"],
		"results":     updateData["results"],
		"task_pid":    updateData["taskPid"],
		"error_code":  updateData["errorCode"],
		"error_msg":   updateData["errorMsg"],
		"exit_code":   updateData["exitCode"],
		"finish_time": updateData["finishTime"],
	}
	WriteResult(updateData)
	return nil
}

func RemoveResult(tid string) error {
	file := fmt.Sprintf("%s/%s.json", entity.ResultDir, tid)
	err := os.Remove(file)
	return err
}

func SyncResult(resultData map[string]string) bool {
	resultData["sign"] = GetSign()
	resultData["tsn"] = "tone20230316-4881"
	tid := resultData["tid"]
	url := fmt.Sprintf("%s/%s", viper.Get("proxy"), entity.SyncResultAPI)
	jsonValue, _ := json.Marshal(resultData)
	client := GetHttpClient()
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Sync results to proxy error, tid:%s | error:%s", tid, err.Error())
		return false
	}
	// SUCCESS 为 "ok" 时同步成功，为 "fail" 时同步失败
	var res map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	err2 := json.Unmarshal(body, &res)
	if err2 != nil {
		return false
	}
	if res["SUCCESS"] == entity.SuccessOk {
		// TODO
		//RemoveResult(tid)
		return true
	}
	return false
}
