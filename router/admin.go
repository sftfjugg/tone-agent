package router

import (
	"github.com/gin-gonic/gin"
	"tone-agent/controller"
)

func AdminRouter(r *gin.Engine) {
	adminRouter := r.Group("")
	adminRouter.GET("/", controller.MainController)
	adminRouter.POST("/config", controller.SetAgentConfig)
}
