package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"tone-agent/entity"
	"tone-agent/core"
)

type TaskController struct {
	beego.Controller
}

func (tc *TaskController) Post() {
	// 接收参数
	task := entity.Task{}
	data := tc.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &task); err == nil {
		// 解析参数失败
	}
	log.Printf("[TaskController]Receive task(tid:%s)", task.Tid)
	response := &entity.AgentResponse{}
	if task.Tid == "" || task.Script == "" {
		log.Println("[TaskController] Task tid or script is null, task cannot running")
		response = &entity.AgentResponse{
			Tid:        task.Tid,
			Success:    "fail",
			TaskStatus: entity.ParamsErrorCode,
			ErrorCode:  entity.ParamsErrorMsg,
		}
	} else {
		response = core.TaskProcessorByPassiveMode(task)
	}
	tc.Data["json"] = response
	tc.ServeJSON()
	tc.StopRun()
}
