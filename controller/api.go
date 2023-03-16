package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"tone-agent/entity"
	"tone-agent/service"
)

func SendTaskController(c *gin.Context) {
	task := entity.Task{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &task)
	log.Printf("[TaskController]Receive task(tid:%s)", task.Tid)
	if task.Sync {
		taskResult := service.ExecTask(task)
		c.JSON(http.StatusOK, gin.H{
			"SUCCESS":     entity.SuccessOk,
			"TID":         task.Tid,
			"STATUS":      entity.TaskCompletedStatus,
			"TASK_RESULT": taskResult["results"],
			"TASK_PID":    taskResult["taskPid"],
			"ERROR_CODE":  taskResult["errorCode"],
			"ERROR_MSG":   taskResult["errorMsg"],
			"FINISH_TIME": taskResult["finishTime"],
		})
	} else {
		go service.ExecTask(task)
		c.JSON(http.StatusOK, gin.H{
			"SUCCESS": entity.SuccessOk,
			"TID":     task.Tid,
			"STATUS":  entity.TaskRunningStatus,
		})
	}
}

func QueryTaskController(c *gin.Context) {
	tid := c.Query("tid")
	taskResult, _ := service.ReadResult(tid)
	if taskResult != nil {
		service.RemoveResult(tid)
		c.JSON(http.StatusOK, gin.H{
			"SUCCESS":     entity.SuccessOk,
			"TID":         tid,
			"STATUS":      taskResult["status"],
			"TASK_RESULT": taskResult["results"],
			"TASK_PID":    taskResult["taskPid"],
			"ERROR_CODE":  taskResult["errorCode"],
			"ERROR_MSG":   taskResult["errorMsg"],
			"FINISH_TIME": taskResult["finishTime"],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"SUCCESS":    entity.SuccessFail,
			"TID":        tid,
			"ERROR_CODE": entity.ResultNotExistErrorCode,
			"ERROR_MSG":  entity.ResultNotExistErrorMsg,
		})
	}

}

func HeartbeatController(c *gin.Context) {
	response := gin.H{
		"SUCCESS": entity.SuccessOk,
		"ARCH":    service.ExecCommand("arch"),
		"KERNEL":  service.ExecCommand("uname -r"),
		"DISTRO":  service.ExecCommand("cat /etc/os-release | grep -i id="),
	}
	c.JSON(http.StatusOK, response)
}
