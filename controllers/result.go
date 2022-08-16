package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"tone-agent/entity"
	"tone-agent/core"
)

type ResultController struct {
	beego.Controller
}

func (rc *ResultController) Get() {
	tid := rc.GetString("tid")
	log.Printf("[ResultController]Query task(tid: %s) request.", tid)
	filename := core.GetFileNameByTid(tid)
	fileExist := core.CheckFileIsExist(filename)
	response := &entity.AgentResponse{
		Tid:     tid,
		Success: "ok",
	}
	if fileExist {
		result := core.ReadFile(filename)
		resultMap := make(map[string]string)
		_ = json.Unmarshal([]byte(result), &resultMap)
		response.TaskStatus = resultMap["status"]
		response.TaskPid = resultMap["taskPid"]
		if resultMap["status"] == entity.TaskCompletedStatus {
			response.ErrorCode = resultMap["errorCode"]
			response.ErrorMsg = resultMap["errorMsg"]
			response.ExitCode = resultMap["exitCode"]
			response.TaskResult = resultMap["result"]
			response.FinishTime = resultMap["finishTime"]
			core.MoveFilePath(filename)
			log.Printf("[ResultController]Task(tid: %s) completed.", tid)
		}
	}
	rc.Data["json"] = response
	rc.ServeJSON()
	rc.StopRun()
}
