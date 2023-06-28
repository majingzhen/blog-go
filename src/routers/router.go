package routers

import (
	commons "blog-go/src/common"
	"blog-go/src/middleware"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

type Routers struct {
	SystemRouter SystemRouter
	commons.Response
}

func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	r.NoRoute(func(c *gin.Context) {
		routers.ErrorMsg(c, "not fount")
		return
	})
	r.NoMethod(func(c *gin.Context) {
		routers.ErrorMsg(c, "not fount")
		return
	})
	// 使用自定义日志配置
	r.Use(middleware.LogMiddleware())
	r.Use(gin.Recovery())

	r.StaticFS("/public", gin.Dir("static/public", true))

	r.Static("/static", "./static/themes")

	var files []string
	filepath.Walk("./static/themes", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".tmpl") {
			files = append(files, path)
		}
		return nil
	})
	r.LoadHTMLFiles(files...)
	// log.Fatalf("Routers InitRouter, values: %v", files)

	sys := r.Group("/")
	{
		sys.GET("/", routers.SystemRouter.Index)
		sys.GET("/login", routers.SystemRouter.Login)
		sys.GET("/captcha", routers.SystemRouter.Captcha)
	}

	// 授权
	// r.Use(middleware.JWTAuthMiddleware())
	r.Use(middleware.StaticPath())
	return r
}
