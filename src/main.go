package main

import (
	"blog-go/src/config"
	"blog-go/src/dao"
	"blog-go/src/server"
)

func init() {
	// 初始化项目配置
	config.InitConfig()
	// 初始化数据库
	dao.InitDataBase()
}
func main() {
	server.Run()
}
