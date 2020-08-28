package test2

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-web/src/boot"
)

type Test2 struct {
}

func NewTest2() *Test2 {
	return &Test2{}
}

func (this *Test2) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(200, "testv2")
	}
}
func (this *Test2) Build(goft *boot.IRoute) {
	goft.Handle("GET", "/test2", this.Get())
}
