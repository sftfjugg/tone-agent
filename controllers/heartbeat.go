package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"tone-agent/core"
)

type HeartbeatController struct {
	beego.Controller
}

func (pc *HeartbeatController) Get() {
	log.Println("[HeartbeatController]heartbeat request from proxy...")
	pc.Data["json"] = map[string]string{
		"SUCCESS": "ok",
		"ARCH": core.ExecCommand("arch"),
		"KERNEL": core.ExecCommand("uname -r"),
		"DISTRO": core.ExecCommand("cat /etc/os-release | grep -i id="),
	}
	pc.ServeJSON()
	pc.StopRun()
}
