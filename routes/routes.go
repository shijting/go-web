package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/controller/test"
	"github.com/shijting/go-web/middlewares"
	"github.com/spf13/viper"
)

func Init() (r *gin.Engine) {
	if viper.GetString("mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.New()

	r.Use(middlewares.Logger(), middlewares.Recovery(true))
	// 路由
	r.POST("/test", test.TestHandle)
	return
}
