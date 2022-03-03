package core

import (
	"github.com/matryer/try"
	"log"
	"time"
	"tone-agent/constant"
)

func TaskProcessorByActiveMode(task constant.Task) {
	//主动模式下执行task
	log.Printf(
		"[TaskProcessorByActiveMode] "+
			"task in process... tid: %s",
		task.Tid,
	)
	var success bool
	err := try.Do(func(attempt int) (bool, error) {
		success = SyncStatusToProxy(task.Tid, "running")
		if attempt > 10 {
			log.Printf(
				"[TaskProcessorByActiveMode]"+
					"sync status to proxy failed, more than 10 retries！tid: %s",
				task.Tid,
			)
			return false, nil
		}
		if !success {
			log.Printf(
				"[TaskProcessorByActiveMode]"+
					"sync status to proxy failed, %d retries！tid: %s",
				attempt,
				task.Tid,
			)
			time.Sleep(time.Duration(2) * time.Second)
			return true, nil
		} else {
			return false, nil
		}
	})
	if err != nil {
		log.Printf("[TaskProcessorByActiveMode]sync status to proxy failed, tid: %s", task.Tid)
	}
	//SyncStatusToProxy(task.Tid, "running")
	if task.Sync {
		// 同步执行任务，然后推送结果
		log.Printf(
			"[TaskProcessorByActiveMode]"+
				"task(tid: %s, sync_type:sync) get ready to exec...",
			task.Tid,
		)
		var startTime = GetCurTimeStr()
		taskPid, result, errorCode, errorMsg, exitCode := ExecCommand(task)
		updateData := map[string]string{
			"tid":           task.Tid,
			"status":        constant.TaskCompletedStatus,
			"task_pid":      taskPid,
			"result":        result,
			"error_code":    errorCode,
			"error_msg":     errorMsg,
			"exit_code":     exitCode,
			"start_time":    startTime,
			"finished_time": GetCurTimeStr(),
		}
		log.Printf(
			"[TaskProcessorByActiveMode]"+
				"task(tid: %s, sync_type:sync) get ready to sync result..., result detail: %s",
			task.Tid,
			updateData,
		)
		err := try.Do(func(attempt int) (bool, error) {
			success = SyncResultToProxy(updateData, true, true)
			if attempt > 10 {
				log.Printf(
					"[TaskProcessorByActiveMode]"+
						"sync exec result to proxy failed, more than 10 retries！tid: %s",
					task.Tid,
				)
				return false, nil
			}
			if !success {
				log.Printf(
					"[TaskProcessorByActiveMode]"+
						"sync exec result to proxy failed, %d retries！tid: %s",
					attempt,
					task.Tid,
				)
				time.Sleep(time.Duration(1) * time.Second)
				return true, nil
			} else {
				log.Printf(
					"[TaskProcessorByActiveMode]"+
						"sync exec result to proxy success！tid: %s",
					task.Tid,
				)
				return false, nil
			}
		})
		if err != nil {
			log.Printf(
				"[TaskProcessorByActiveMode]"+
					"sync status to proxy failed, tid: %s",
				task.Tid,
			)
		}
		//SyncResultToProxy(updateData, true, true)
	} else {
		log.Printf(
			"[TaskProcessorByActiveMode]"+
				"task(tid: %s, sync_type:async) get ready to exec...",
			task.Tid,
		)
		go ExecCommand(task)
		log.Printf("[TaskProcessorByActiveMode]task(tid: %s) exec...", task.Tid)
	}
}

func TaskProcessorByPassiveMode(task constant.Task) *constant.AgentResponse {
	// 被动模式下执行task
	response := &constant.AgentResponse{}
	if task.Sync {
		taskPid, result, errorCode, errorMsg, exitCode := ExecCommand(task)
		//filename := GetFileNameByTid(task.Tid)
		//MoveFilePath(filename)
		log.Printf("[TaskController]task(tid: %s) completed.", task.Tid)
		response = &constant.AgentResponse{
			Tid:        task.Tid,
			Success:    "ok",
			TaskStatus: "completed",
			TaskResult: result,
			TaskPid:    taskPid,
			ErrorCode:  errorCode,
			ErrorMsg:   errorMsg,
			ExitCode:   exitCode,
		}
	} else {
		go ExecCommand(task)
		response = &constant.AgentResponse{
			Tid:        task.Tid,
			Success:    "ok",
			TaskStatus: "running",
		}
		log.Printf("[TaskController]task(tid: %s) running.", task.Tid)
	}
	return response
}
