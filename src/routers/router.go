package routers

import (
	"blog-go/src/middleware"
	"github.com/gin-gonic/gin"
)

type Routers struct {
}

func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	// 使用自定义日志配置
	r.Use(middleware.LogMiddleware())
	r.Use(gin.Recovery())
	return r
}
