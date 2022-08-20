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
	arch, _ := core.ExecCommand("arch")
	kernel, _ := core.ExecCommand("uname -r")
	distro, _ := core.ExecCommand("cat /etc/os-release | grep -i id=")
	hbc.Data["json"] = map[string]string{
		"SUCCESS": "ok",
		"ARCH":    arch,
		"KERNEL":  kernel,
		"DISTRO":  distro,
	}
	hbc.ServeJSON()
	hbc.StopRun()
}
