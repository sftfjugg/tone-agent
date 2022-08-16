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

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (gtc *GenerateTSNController) Get() {
	macAddrs := comm.GetMacAddrs()
	macAddrStr := strings.Join(macAddrs, `|`)
	macAddrEncrypt := comm.MD5Encrypt(macAddrStr, "tone-agent")
	response := &entity.TSNResponse{
		Code: 200,
		Msg:  "success",
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

	response := &entity.TSNResponse{
		Code: 200,
		Msg:  "success",
		TSN:  viper.GetString("tsn"),
		Mode:  viper.GetString("mode"),
		Proxy:  viper.GetString("proxy"),
	}
	scc.Data["json"] = response
	scc.ServeJSON()
	scc.StopRun()
}

func (gcc *GetConfigController) Get() {
	//var conf entity.Config
	//conf.GetConf()
	response := &entity.TSNResponse{
		Code: 200,
		Msg:  "success",
		TSN:  viper.GetString("tsn"),
		Mode:  viper.GetString("mode"),
		Proxy:  viper.GetString("proxy"),
	}
	gcc.Data["json"] = response
	gcc.ServeJSON()
	gcc.StopRun()
}

func (giac *GetIpAddrController) Get() {
	ipAddr := comm.GetLocalIP()
	response := &entity.IPResponse{
		Code: 200,
		Msg:  "success",
		IP:   ipAddr,
	}
	giac.Data["json"] = response
	giac.ServeJSON()
	giac.StopRun()
}
