package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type HeartbeatController struct {
	beego.Controller
}

func (pc *HeartbeatController) Get() {
	log.Println("[HeartbeatController]heartbeat request from proxy...")
	pc.Data["json"] = map[string]string{"SUCCESS": "ok"}
	pc.ServeJSON()
	pc.StopRun()
}
