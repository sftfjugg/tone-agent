package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/astaxie/beego"
	"github.com/matryer/try"
	"github.com/spf13/viper"

	"tone-agent/entity"
)

func decodeScript(task entity.Task) string {
	uDec, err := base64.StdEncoding.DecodeString(task.Script)
	if err != nil {
		log.Printf(
			"[ExecTask]task(tid: %s) script decode error failed error! error: %s",
			task.Tid,
			err,
		)
	}
	script := string(uDec)
	return script
}

func getExecCmd(script string, task entity.Task) string {
	var execCmd string
	if task.ScriptType == "cmd" {
		execCmd = fmt.Sprintf("%s %s", script, task.Args)
	} else {
		// 将脚本保存成一个临时文件
		tmpScriptDir := beego.AppConfig.String("TmpScriptFileDir")
		if !CheckFileIsExist(tmpScriptDir) {
			MakeDir(tmpScriptDir)
		}
		tmpShellName := fmt.Sprintf(
			"%s/tmp_script_%s.sh",
			tmpScriptDir,
			RandStringBytesMaskImprSrcUnsafe(16),
		)
		suc, _ := WriteFile(tmpShellName, []byte(script))
		execCmd = fmt.Sprintf("sh %s %s", tmpShellName, task.Args)
		log.Printf(
			"[ExecTask]task(tid: %s) script write to a temp file.success(%s)",
			task.Tid,
			strconv.FormatBool(suc),
		)
	}
	return execCmd
}

func getEnv(task entity.Task) []string {
	envArr := strings.Split(task.Env, ",")
	// 	for _, e := range os.Environ() {
	//         envArr = append(envArr, e)
	//     }
	//     log.Printf("env:%s", envArr)
	envArr = append(envArr, fmt.Sprintf("PATH=%s", os.Getenv("PATH")))
	envArr = append(envArr, "HOME=/root")
	return envArr
}

func syncExecInfo(task entity.Task, taskPid int, script string) {
	runningResultMap := map[string]string{}
	if taskPid != 0 && strings.Contains(script, "reboot") {
		runningResultMap = map[string]string{
			"status":   entity.TaskCompletedStatus,
			"tid":      task.Tid,
			"task_pid": strconv.Itoa(taskPid),
			//"script": script,
			"result":      "reboot success",
			"start_time":  GetCurTimeStr(),
			"finish_time": GetCurTimeStr(),
		}
		log.Printf(
			"[syncExecInfo]the reboot command executed success...server begin restart. tid: %s",
			task.Tid,
		)
		if viper.Get("mode") == "active" {
			// active模式下执行reboot命令，将重启执行结果同步到proxy端
			updateDate := map[string]string{
				"status":      entity.TaskCompletedStatus,
				"tid":         task.Tid,
				"task_pid":    strconv.Itoa(taskPid),
				"result":      "reboot command executed success",
				"start_time":  GetCurTimeStr(),
				"finish_time": GetCurTimeStr(),
			}
			var success bool
			err := try.Do(func(attempt int) (bool, error) {
				success = SyncResultToProxy(updateDate, false, true)
				if attempt > 10 {
					log.Printf(
						"[syncExecInfo]sync exec info to proxy failed, "+
							"more than 10 retries！tid: %s",
						task.Tid,
					)
					return false, nil
				}
				if !success {
					log.Printf(
						"[syncExecInfo]sync exec info to proxy failed,%d retries！tid: %s",
						attempt,
						task.Tid,
					)
					time.Sleep(time.Duration(1) * time.Second)
					return true, nil
				} else {
					log.Printf(
						"[syncExecInfo]sync exec info to proxy success！tid: %s",
						task.Tid,
					)
					return false, nil
				}
			})
			if err != nil {
				log.Printf(
					"[syncExecInfo]sync status to proxy failed, tid: %s",
					task.Tid,
				)
			}
			//SyncResultToProxy(updateDate, false, true)
			WriteResult(runningResultMap, true)
		} else {
			// passive模式下的reboot命令，proxy端已有结果，无需同步
			WriteResult(runningResultMap, true)
		}
	} else {
		runningResultMap = map[string]string{
			"status":  entity.TaskRunningStatus,
			"tid":     task.Tid,
			"taskPid": strconv.Itoa(taskPid),
			//"script": script,
			"startTime": GetCurTimeStr(),
		}
		if viper.Get("mode") == "active" {
			var success bool
			err := try.Do(func(attempt int) (bool, error) {
				success = SyncExecTimeToProxy(
					task.Tid,
					"start_time",
					strconv.Itoa(taskPid),
				)
				if attempt > 10 {
					log.Printf(
						"[syncExecInfo]sync exec time to proxy failed,"+
							"more than 10 retries！tid: %s",
						task.Tid,
					)
					return false, nil
				}
				if !success {
					log.Printf(
						"[syncExecInfo]"+
							"sync exec time to proxy failed, %d retries！tid: %s",
						attempt,
						task.Tid,
					)
					time.Sleep(time.Duration(1) * time.Second)
					return true, nil
				} else {
					log.Printf(
						"[syncExecInfo]sync exec time to proxy success. tid: %s",
						task.Tid,
					)
					return false, nil
				}
			})
			if err != nil {
				log.Printf(
					"[TaskProcessorByActiveMode]sync status to proxy failed, tid: %s",
					task.Tid,
				)
			}
			//SyncExecTimeToProxy(task.Tid, "start_time")
		}
		log.Printf("[syncExecInfo] sync exec info end. tid: %s", task.Tid)
		WriteResult(runningResultMap, task.Sync)
	}
}

func getExecErrInfo(execErr error, task entity.Task) (string, string, string) {
	var exitCode string
	errorCode := entity.ExecCmdErrorCode
	errorMsg := fmt.Sprintf("%s(%s)", entity.ExecCmdErrorMsg, execErr.Error())

	if exiterr, ok := execErr.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			exitCode = strconv.Itoa(status.ExitStatus())
		}
	}
	log.Printf(
		"[ExecTask]task(tid: %s) script exec failed! detail: %s",
		task.Tid,
		execErr.Error(),
	)
	return errorCode, errorMsg, exitCode
}

func syncErrorResult(tid string, errorCode string, errorMsg string, sync bool) {
	resultMap := map[string]string{
		"tid":        tid,
		"status":     entity.TaskCompletedStatus,
		"errorCode":  errorCode,
		"errorMsg":   errorMsg,
		"finishTime": GetCurTimeStr(),
	}
	WriteResult(resultMap, sync)
}

func writeExecResult(
	task entity.Task,
	taskPid int,
	errorCode string,
	errorMsg string,
	result string,
	exitCode string) {
	resultMap := map[string]string{
		"status": entity.TaskCompletedStatus,
		"tid":    task.Tid,
		//"script":     task.Script,
		"taskPid":    strconv.Itoa(taskPid),
		"errorCode":  errorCode,
		"errorMsg":   errorMsg,
		"result":     result,
		"exitCode":   exitCode,
		"finishTime": GetCurTimeStr(),
	}
	log.Printf(
		"[ExecTask]"+
			"task(tid: %s) begin to write exec result. errorCode:%s | errorMsg:%s",
		task.Tid,
		errorCode,
		errorMsg,
	)
	WriteResult(resultMap, task.Sync)
	log.Printf(
		"[ExecTask]"+
			"task(tid: %s) write exec result completed. errorCode:%s | errorMsg:%s",
		task.Tid,
		errorCode,
		errorMsg,
	)
}

func ExecTask(task entity.Task) (string, string, string, string, string) {
	script := decodeScript(task)
	execCmd := getExecCmd(script, task)
	env := getEnv(task)

	ctxt, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(task.Timeout)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctxt, "/bin/bash", "-c", execCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pgid: 0}
	cmd.Dir = task.Cwd
	cmd.Env = env

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	cmdErr := cmd.Start()
	if cmdErr != nil {
		log.Printf(
			"[ExecTask]cmd start failed!, error:%s",
			cmdErr.Error(),
		)
		syncErrorResult(
			task.Tid,
			"a-1001",
			"cmd start failed",
			task.Sync,
		)
	}
	log.Printf(""+
		"[ExecTask]cmd start...., tid:%s | sync_type:%t",
		task.Tid,
		task.Sync,
	)
	taskPid := cmd.Process.Pid
	log.Printf(
		"[ExecTask]cmd process pid is %d | tid:%s | sync_type:%t",
		taskPid,
		task.Tid,
		task.Sync,
	)
	if task.Sync {
		log.Printf(
			"[ExecTask]tid:%s | sync_type:sync without sync exec info",
			task.Tid,
		)
	} else {
		log.Printf(
			"[ExecTask]tid:%s | sync_type:async begin to sync exec info",
			task.Tid,
		)
		syncExecInfo(task, taskPid, script)
	}
	var (
		errorCode string
		errorMsg  string
		exitCode  string
	)
	if execErr := cmd.Wait(); execErr != nil {
		errorCode, errorMsg, exitCode = getExecErrInfo(execErr, task)
	}
	stdout := b.Bytes()
	result := string(stdout)

	writeExecResult(task, taskPid, errorCode, errorMsg, result, exitCode)
	return strconv.Itoa(taskPid), result, errorCode, errorMsg, exitCode
}

func ExecCommand(execCmd string) (string, ) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", execCmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		log.Printf("[ExecCommand]cmd run failed: %s | detail:%s \n", err, errStr)
		return ""
	}
	return outStr
}
