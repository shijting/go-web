package test

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/src/boot"
	"github.com/spf13/viper"
)

type TestClass struct {
}

func NewTestClass() *TestClass {
	return &TestClass{}
}

type Params struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func (this *TestClass) TestHandle() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 参数验证
		//var p = &Params{}
		//if err := c.ShouldBind(p); err != nil {
		//	fmt.Printf("%v\n", err)
		//	// 请求参数有误，直接返回响应
		//	resp := common.ValidateError(err)
		//	c.JSON(200, resp)
		//	return
		//}
		//fmt.Printf("%#v\n", p)
		c.String(200, viper.GetString("mysql.host"))
	}

}

func (this *TestClass) Build(goft *boot.IRoute) {
	goft.Handle("GET", "/test", this.TestHandle())
}
