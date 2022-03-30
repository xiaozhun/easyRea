package initialize

import (
	"easyReq/api"
	"easyReq/global"
	"easyReq/middleware"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	PublicGroup := Router.Group("")
	{
		// 初始化配置路由
		PublicGroup.Any(global.GVA_PATH, api.Index)
		PublicGroup.POST("noteSmsRecord", api.NoteSmsRecord)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
