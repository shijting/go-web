package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/common"
	"github.com/spf13/viper"
)

type Params struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func TestHandle(c *gin.Context) {
	// 操作redis
	//err := redis.GetRedisInstance().Set("score", 100, 0).Err()
	//if err != nil {
	//	fmt.Printf("set score failed, err:%v\n", err)
	//	return
	//}
	// 参数验证
	var p = &Params{}
	if err := c.ShouldBind(p); err != nil {
		fmt.Printf("%v\n", err)
		// 请求参数有误，直接返回响应
		resp := common.ValidateError(err)
		c.JSON(200, resp)
		return
	}
	fmt.Printf("%#v\n", p)
	c.String(200, viper.GetString("mysql.host"))
}
