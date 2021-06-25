package main

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go_gin/global"
	"go_gin/initialize"
	"time"
)

func main() {
	//1.初始化yaml配置
	initialize.InitConfig()
	//2. 初始化routers
	Router := initialize.Routers()
	// 3.初始化日志信息R
	initialize.InitLogger()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// 5.初始化mysql
	initialize.InitMysqlDB()
	//6. 初始化redis
	initialize.InitRedis()
	// 7. 初始化minIO
	initialize.InitMinIO()
	color.Yellow(">>>>>>>>>>>>>>>>>>gin服务开始了~" )
	global.Redis.Set("test", "testValue", time.Second)
	//time.Sleep(time.Second*2)
	value := global.Redis.Get("test")
	color.Blue(value.Val())
	// 设置启动端口
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	//
	if err != nil {
		zap.L().Info("this is hello func", zap.String("user", "xixixi"))
	}
}
