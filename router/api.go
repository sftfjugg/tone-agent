package router

import (
	"github.com/gin-gonic/gin"
	"tone-agent/controller"
)

func APIRouter(r *gin.Engine) {
	apiRouter := r.Group("api")
	apiRouter.POST("/task", controller.SendTaskController)
	apiRouter.GET("/query", controller.QueryTaskController)
	apiRouter.GET("/heartbeat", controller.HeartbeatController)
}
