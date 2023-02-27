package main

import (
	"flag"
	"wechatbot/bootstrap"
	"wechatbot/config"
)

func main() {
	// 读取全局配置文件
	cfg := flag.String("c", "", "configuration file")
	flag.Parse()
	config.LoadConfig(*cfg)

	bootstrap.Run()
}
