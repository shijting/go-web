package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/src/boot"
	"github.com/shijting/go-web/src/controller/test"
	"github.com/shijting/go-web/src/controller/test2"
	"github.com/shijting/go-web/src/middlewares"
	"github.com/spf13/viper"
)

func Init() (r *boot.IRoute) {
	if viper.GetString("mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r = boot.NewIRoute()
	// 加载中间件
	r.Use(middlewares.Logger(), middlewares.Recovery(true))
	// 路由
	r.Mount("v1", test.NewTestClass()). // http://localhost:3000/v1/test
						Mount("v2", test2.NewTest2(), test.NewTestClass()) //http://localhost:3000/v2/test2
	return
}
