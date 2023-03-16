package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
	"tone-agent/entity"
)

func PullTask() []entity.Task {
	url := fmt.Sprintf("%s/%s", viper.Get("proxy"), entity.PullTaskAPI)
	tsn := viper.GetString("tsn")
	sign := GetSign()
	data := map[string]string{"tsn": tsn, "sign": sign}
	jsonData, _ := json.Marshal(data)
	client := GetHttpClient()
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != 200 {
		//TODO
		log.Println(fmt.Sprintf("Pull task failed.url:%s | error:%s", url, err))
		return nil
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	taskRes := &entity.PullTaskRes{}
	json.Unmarshal(result, taskRes)
	return taskRes.Tasks
}

func ProcessTask(tasks []entity.Task) {
	for _, task := range tasks {
		ExecTask(task)
	}
}

func ExecTask(task entity.Task) map[string]string {
	uDec, _ := base64.StdEncoding.DecodeString(task.Script)
	script := string(uDec)

	env := strings.Split(task.Env, ",")
	env = append(env, fmt.Sprintf("PATH=%s", os.Getenv("PATH")))
	env = append(env, "HOME=/root")

	ctxt, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(task.Timeout)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctxt, "/bin/bash", "-c", script)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pgid: 0}
	cmd.Dir = task.Cwd
	cmd.Env = env

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	cmd.Start()

	taskPid := cmd.Process.Pid
	var (
		errorCode string
		errorMsg  string
		exitCode  string
	)
	if execErr := cmd.Wait(); execErr != nil {
		errorCode = entity.ExecCmdErrorCode
		errorMsg = fmt.Sprintf("%s(%s)", entity.ExecCmdErrorMsg, execErr.Error())

		if exitErr, ok := execErr.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = strconv.Itoa(status.ExitStatus())
			}
		}
	}
	stdout := b.Bytes()
	result := string(stdout)
	resultData := map[string]string{
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
	WriteResult(resultData)
	return resultData
}

func ExecCommand(execCmd string) string {
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
