package routers

import "github.com/astaxie/beego"
import "tone-agent/controllers"


func init() {
	// api
	beego.Router("api/task", &controllers.TaskController{})
	beego.Router("api/query", &controllers.ResultController{})
	beego.Router("api/heartbeat", &controllers.HeartbeatController{})

	// index
	beego.Router("/", &controllers.MainController{})
	beego.Router("tsn/generate", &controllers.GenerateTSNController{})
	beego.Router("ip/info", &controllers.GetIpAddrController{})
}