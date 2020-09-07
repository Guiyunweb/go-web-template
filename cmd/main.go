package main

import (
	"github.com/Guiyunweb/go-web-template/conf"
	"github.com/Guiyunweb/go-web-template/library/log"
	"github.com/Guiyunweb/go-web-template/server"
)

func main() {
	// 读取解析配置文件
	if err := conf.Init(); err != nil {
		log.Panic("初始化配置失败")
	}
	server.Run()
}
