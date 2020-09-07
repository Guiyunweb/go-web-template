package main

import (
	"github.com/Guiyunweb/go-web-template/conf"
	"github.com/Guiyunweb/go-web-template/library/log"
)

func main() {
	// 读取解析配置文件
	if err := conf.Init(); err != nil {
		log.Panic("初始化配置失败")
	}
	log.Info("运行成功")
}
