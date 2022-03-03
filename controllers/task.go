package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"tone-agent/constant"
	"tone-agent/core"
)

type TaskController struct {
	beego.Controller
}


func (pc *TaskController) Post() {
	// 接收参数
	task := constant.Task{}
	data := pc.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &task); err == nil {
		// 解析参数失败
	}
	log.Printf(
		"[TaskController]Receive task(tid:%s) | sync:%s | scriptType:%s | args:%s | env:%s | cwd:%s",
		task.Tid,
		strconv.FormatBool(task.Sync),
		task.ScriptType,
		task.Args,
		task.Env,
		task.Cwd,
	)
	response := &constant.AgentResponse{}
	if task.Tid == "" || task.Script == ""{
		log.Println("[TaskController] Task tid or script is null, task cannot running")
		response = &constant.AgentResponse{
			Tid:        task.Tid,
			Success:    "fail",
			TaskStatus: constant.ParamsErrorCode,
			ErrorCode: constant.ParamsErrorMsg,
		}
	}else{
		response = core.TaskProcessorByPassiveMode(task)
	}
	pc.Data["json"] = response
	pc.ServeJSON()
	pc.StopRun()
}