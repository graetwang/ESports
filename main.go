package main

import (
	"flag"
	"go.uber.org/zap"
	"server/cache/redis"
	"server/conf"
	"server/dao/mysql"
	"server/logger"
	"server/router"
)

func main() {
	var cfn string
	// 0.从命令行获取可能的conf路径
	// goods_service -conf="./conf/config_qa.yaml"
	// goods_service -conf="./conf/config_online.yaml"
	flag.StringVar(&cfn, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()
	// 1. 读取并解析yaml配置文件
	err := conf.Init(cfn)
	if err != nil {
		zap.L().Error("配置文件读取失败")
		return
	}
	//2.初始化日志
	err = logger.Init(conf.Conf.LogConfig, conf.Conf.Mode)
	if err != nil {
		zap.L().Error("日志初始化失败")
		return
	}
	//3.初始化msql
	err = mysql.Init(conf.Conf.MySQLConfig)
	if err != nil {
		zap.L().Error("mysql初始化失败")
		return
	}
	//4.初始化redis
	err = redis.Init(conf.Conf.RedisConfig)
	if err != nil {
		zap.L().Error("redis初始化失败")
		return
	}
	// 6. 注册gin路由
	// 以下代码需要在 main goroutine 里执行
	// 保证 web server可以正常关闭
	// 以下代码需要在 main goroutine 里执行
	// 保证
	r := router.NewRouter()
	_ = r.Run(conf.Conf.Port)
}
