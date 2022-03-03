package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"tone-agent/constant"
	"tone-agent/core"
)

type ResultController struct {
	beego.Controller
}

func (pc *ResultController) Get() {
	tid := pc.GetString("tid")
	log.Printf("[ResultController]Query task(tid: %s) request.", tid)
	filename := core.GetFileNameByTid(tid)
	fileExist := core.CheckFileIsExist(filename)
	response := &constant.AgentResponse{
		Tid:     tid,
		Success: "ok",
	}
	if fileExist {
		result := core.ReadFile(filename)
		resultMap := make(map[string]string)
		_ = json.Unmarshal([]byte(result), &resultMap)
		response.TaskStatus = resultMap["status"]
		response.TaskPid = resultMap["taskPid"]
		if resultMap["status"] == constant.TaskCompletedStatus {
			response.ErrorCode = resultMap["errorCode"]
			response.ErrorMsg = resultMap["errorMsg"]
			response.ExitCode = resultMap["exitCode"]
			response.TaskResult = resultMap["result"]
			response.FinishTime = resultMap["finishTime"]
			core.MoveFilePath(filename)
			log.Printf("[ResultController]Task(tid: %s) completed.", tid)
		}
	}
	pc.Data["json"] = response
	pc.ServeJSON()
	pc.StopRun()
}
