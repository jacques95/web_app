package main

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logger"
	"web_app/settings"
)

func main() {
	// 1、 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 2、 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success.")

	// 3、 初始化 MySQL 连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4、 初始化 Redis 连接
	// 5、 注册路由器
	// 6、 启动服务

}
