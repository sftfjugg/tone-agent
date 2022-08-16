package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
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

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (gtc *GenerateTSNController) Get() {
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
