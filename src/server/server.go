package server

import (
	"blog-go/src/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)

func Run() {
	gin.SetMode(viper.GetString("app.runMode"))
	// 关闭日志颜色
	gin.DisableConsoleColor()
	log, _ := os.Create(viper.GetString("app.logFile"))
	gin.DefaultWriter = io.MultiWriter(log)
	router := new(routers.Routers).InitRouter()

	router.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
}
