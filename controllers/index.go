package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"strings"
	"tone-agent/comm"
	"tone-agent/core"
	"tone-agent/entity"
)

type MainController struct {
	beego.Controller
}

type GenerateTSNController struct {
	beego.Controller
}

type SetConfigController struct {
	beego.Controller
}

type GetConfigController struct {
	beego.Controller
}

type GetIpAddrController struct {
	beego.Controller
}

type GetLogController struct {
	beego.Controller
}

type RestartServiceController struct {
	beego.Controller
}

type StopServiceController struct {
	beego.Controller
}

type SendHeartbeatController struct {
	beego.Controller
}

type RequestTestController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Render()
}

func (gtc *GenerateTSNController) Post() {
	macAddrs := comm.GetMacAddrs()
	macAddrStr := strings.Join(macAddrs, `|`)
	macAddrEncrypt := comm.MD5Encrypt(macAddrStr, entity.AgentTSNSalt)
	response := &entity.TSNResponse{
		Code: entity.SuccessCode,
		Msg:  entity.SuccessMsg,
		TSN:  macAddrEncrypt,
	}
	gtc.Data["json"] = response
	gtc.ServeJSON()
	gtc.StopRun()
}

func (scc *SetConfigController) Post() {
	config := entity.Config{}
	data := scc.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &config); err != nil {
		response := &entity.ErrorResponse{
			Code: entity.PidNotExistErrorCode,
			Msg:  entity.ParamsErrorMsg,
		}
		scc.Data["json"] = response
		scc.ServeJSON()
		scc.StopRun()
	}
	// 修改配置
	err := comm.SetConfig(
			config.TSN,
			config.Mode,
			config.Proxy,
		)
	if err != nil{
		panic(err)
	}
	response := &entity.ConfigResponse{
		Code: entity.SuccessCode,
		Msg:  entity.SuccessMsg,
		Config: entity.Config{
			TSN:  viper.GetString("tsn"),
			Mode:  viper.GetString("mode"),
			Proxy:  viper.GetString("proxy"),
		},
	}
	scc.Data["json"] = response
	scc.ServeJSON()
	scc.StopRun()
}

func (gcc *GetConfigController) Get() {
	response := &entity.ConfigResponse{
		Code: entity.SuccessCode,
		Msg:  entity.SuccessMsg,
		Config: entity.Config{
			TSN:  viper.GetString("tsn"),
			Mode:  viper.GetString("mode"),
			Proxy:  viper.GetString("proxy"),
		},
	}
	gcc.Data["json"] = response
	gcc.ServeJSON()
	gcc.StopRun()
}

func (gcc *GetLogController) Get() {
	log := comm.GetLog()
	response := &entity.LogResponse{
		Code: entity.SuccessCode,
		Msg:  entity.SuccessMsg,
		Log: log,
	}
	gcc.Data["json"] = response
	gcc.ServeJSON()
	gcc.StopRun()
}

func (giac *GetIpAddrController) Get() {
	IPAddr := comm.GetLocalIP()
	response := &entity.IPResponse{
		Code: entity.SuccessCode,
		Msg:  entity.SuccessMsg,
		IP:   IPAddr,
	}
	giac.Data["json"] = response
	giac.ServeJSON()
	giac.StopRun()
}

func (rsc *RestartServiceController) Post() {
	result, err := core.ExecCommand("systemctl restart toneagent")
	if err != ""{
		response := &entity.ErrorResponse{
			Code: entity.ExecCmdErrorCode,
			Msg:  err,
		}
		rsc.Data["json"] = response
		rsc.ServeJSON()
		rsc.StopRun()
	}else{
		response := &entity.SuccessResponse{
			Code: entity.SuccessCode,
			Msg:  result,
		}
		rsc.Data["json"] = response
		rsc.ServeJSON()
		rsc.StopRun()
	}
}

func (ssc *StopServiceController) Post() {
	result, err := core.ExecCommand("systemctl stop toneagent")
	if err != ""{
		response := &entity.ErrorResponse{
			Code: entity.ExecCmdErrorCode,
			Msg:  err,
		}
		ssc.Data["json"] = response
		ssc.ServeJSON()
		ssc.StopRun()
	}else{
		response := &entity.SuccessResponse{
			Code: entity.SuccessCode,
			Msg:  result,
		}
		ssc.Data["json"] = response
		ssc.ServeJSON()
		ssc.StopRun()
	}
}

func (shc *SendHeartbeatController) Post() {
	err := core.SyncHeartbeatToProxy()
	if err != ""{
		response := &entity.ErrorResponse{
			Code: entity.SyncHeartbeatErrorCode,
			Msg:  err,
		}
		shc.Data["json"] = response
		shc.ServeJSON()
		shc.StopRun()
	}else{
		response := &entity.SuccessResponse{
			Code: entity.SuccessCode,
			Msg:  entity.SuccessMsg,
		}
		shc.Data["json"] = response
		shc.ServeJSON()
		shc.StopRun()
	}
}

func (rtc *RequestTestController) Get() {
	client := core.GetHttpClient()
	url := rtc.GetString("url")
	resp, _ := client.Get(url)
	if resp.StatusCode != 200 {
		response := &entity.ErrorResponse{
			Code: entity.RequestErrorCode,
			Msg:  entity.RequestErrorMsg,
		}
		rtc.Data["json"] = response
		rtc.ServeJSON()
		rtc.StopRun()
	}else{
		response := &entity.SuccessResponse{
			Code: entity.SuccessCode,
			Msg:  entity.SuccessMsg,
		}
		rtc.Data["json"] = response
		rtc.ServeJSON()
		rtc.StopRun()
	}
}