package setup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/boot/redis"
	"github.com/shijting/go-web/middlewares"
	"github.com/spf13/viper"
)

func Init() (r *gin.Engine) {
	r = gin.New()
	r.Use(middlewares.Logger(), middlewares.Recovery(true))
	r.GET("/", func(context *gin.Context) {
		err := redis.GetRedisInstance().Set("score", 100, 0).Err()
		if err != nil {
			fmt.Printf("set score failed, err:%v\n", err)
			return
		}
		context.String(200, viper.GetString("mysql.host"))
	})
	return
}
