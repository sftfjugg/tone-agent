package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"tone-agent/comm"
	"tone-agent/entity"
)

type MainController struct {
	beego.Controller
}

type GenerateTSNController struct {
	beego.Controller
}

type GetIpAddrController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (gtc *GenerateTSNController) Get() {
	macAddrs := comm.GetMacAddrs()
	macAddrStr := strings.Join(macAddrs, `|`)
	macAddrEncrypt := comm.MD5Encrypt(macAddrStr, "tone-agent")
	response := &entity.TSNResponse{
		CODE:	200,
		MSG:	"success",
		TSN:	macAddrEncrypt,
	}
	gtc.Data["json"] = response
	gtc.ServeJSON()
	gtc.StopRun()
}

func (giac *GetIpAddrController) Get() {
	ipAddr := comm.GetLocalIP()
	response := &entity.IPResponse{
		CODE:	200,
		MSG:	"success",
		IP:	ipAddr,
	}
	giac.Data["json"] = response
	giac.ServeJSON()
	giac.StopRun()
}