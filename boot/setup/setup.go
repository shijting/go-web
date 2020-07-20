package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/middlewares"
	"github.com/spf13/viper"
)

func Init() (r *gin.Engine) {
	r = gin.New()
	r.Use(middlewares.Logger(), middlewares.Recovery(true))
	r.GET("/", func(context *gin.Context) {

		context.String(200, viper.GetString("mysql.host"))
	})
	return
}
