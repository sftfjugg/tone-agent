package controllers

import (
	"log"

	"github.com/astaxie/beego"

	"tone-agent/core"
)

type HeartbeatController struct {
	beego.Controller
}

func (hbc *HeartbeatController) Get() {
	log.Println("[HeartbeatController]heartbeat request from proxy...")
	hbc.Data["json"] = map[string]string{
		"SUCCESS": "ok",
		"ARCH":    core.ExecCommand("arch"),
		"KERNEL":  core.ExecCommand("uname -r"),
		"DISTRO":  core.ExecCommand("cat /etc/os-release | grep -i id="),
	}
	hbc.ServeJSON()
	hbc.StopRun()
}
