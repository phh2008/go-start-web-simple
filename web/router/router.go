package router

import (
	"com.gientech/equipment-data-collection/app"
	"github.com/gin-gonic/gin"
)

// Register 注册路由
func Register(app *gin.Engine, appCtx *app.AppContext) {

	// auth
	v1 := app.Group("/api/v1")
	// 中间件
	// v1.Use(x)
	{
		v1.GET("/hello", appCtx.HelloApi.Hello)
	}

}
