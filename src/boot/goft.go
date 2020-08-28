package boot

import (
	"github.com/gin-gonic/gin"
)

type IRoute struct {
	*gin.Engine
	group *gin.RouterGroup
}

func NewIRoute() *IRoute {
	return &IRoute{Engine: gin.New()}
}

func (this *IRoute) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *IRoute {
	this.group.Handle(httpMethod, relativePath, handlers...)
	return this
}

// 挂载
func (this *IRoute) Mount(group string, classes ...IClass) *IRoute {
	this.group = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
