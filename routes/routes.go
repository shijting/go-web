package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/controller/test"
	"github.com/shijting/go-web/middlewares"
)

func Init() (r *gin.Engine) {
	r = gin.New()
	r.Use(middlewares.Logger(), middlewares.Recovery(true))
	// 路由
	r.POST("/test", test.TestHandle)
	return
}
